package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type StoreDataStatusDeviceStateRequest struct {
	UptimeSeconds int `json:"uptimeSeconds"`
}

type StoreDataStatusDeviceInfoRequest struct {
	HardwareVersion string `json:"hardwareVersion"`
	SoftwareVersion string `json:"softwareVersion"`
}

type StoreDataStatusRequest struct {
	DeviceInfo            StoreDataStatusDeviceInfoRequest  `json:"deviceInfo"`
	DeviceState           StoreDataStatusDeviceStateRequest `json:"deviceState"`
	State                 string                            `json:"state"`
	SNR                   float64                           `json:"snr"`
	DownlinkThroughputBps float64                           `json:"downlinkThroughputBps"`
	UplinkThroughputBps   float64                           `json:"uplinkThroughputBps"`
	PopPingLatencyMs      float64                           `json:"popPingLatencyMs"`
	PopPingDropRate       float64                           `json:"popPingDropRate"`
	PercentObstructed     float64                           `json:"percentObstructed"`
	SecondsObstructed     float64                           `json:"secondsObstructed"`
}

type StoreDataRequest struct {
	When   string                 `json:"when"`
	Status StoreDataStatusRequest `json:"status"`
}

type StoreDataResponse struct {
	Error string `json:"error,omitempty"`
}

func StoreData(w http.ResponseWriter, r *http.Request) {
	storeDataResponse := StoreDataResponse{}

	storeDataRequest := StoreDataRequest{}
	if err := json.NewDecoder(r.Body).Decode(&storeDataRequest); err != nil {
		logger.Error(err)
		storeDataResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, storeDataResponse)
		return
	}

	// update the "last received data from" date of the dish
	dishy := DishyFromTokenContext(r)
	if dishy == nil {
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	when, err := time.Parse(time.RFC3339, storeDataRequest.When)
	if err != nil {
		logger.Error(err)
		storeDataResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, storeDataResponse)
		return
	}

	downlinkThroughputBps := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "downlink_throughput_bps",
		Help: "The reported downlink throughput in bits per second.",
	})
	downlinkThroughputBps.Set(storeDataRequest.Status.DownlinkThroughputBps)

	uplinkThroughputBps := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "uplink_throughput_bps",
		Help: "The reported uplink throughput in bits per second.",
	})
	uplinkThroughputBps.Set(storeDataRequest.Status.UplinkThroughputBps)

	snr := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "snr",
		Help: "The reported signal to noise ratio.",
	})
	snr.Set(storeDataRequest.Status.SNR)

	popPingLatencyMs := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pop_ping_latency_ms",
		Help: "The reported ping latency to the POP.",
	})
	popPingLatencyMs.Set(storeDataRequest.Status.PopPingLatencyMs)

	popPingDropRate := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pop_ping_drop_rate",
		Help: "The reported drop rate of the ping to the POP.",
	})
	popPingDropRate.Set(storeDataRequest.Status.PopPingDropRate)

	percentObstructed := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "percent_obstructed",
		Help: "The reported percent obstructed.",
	})
	percentObstructed.Set(storeDataRequest.Status.PercentObstructed)

	secondsObstructed := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "seconds_obstructed",
		Help: "The reported seconds obstructed.",
	})
	secondsObstructed.Set(storeDataRequest.Status.SecondsObstructed)

	registry := prometheus.NewRegistry()

	pusher := push.New("http://prom-pushgateway:9091", dishy.ID).Gatherer(registry)
	pusher.Collector(downlinkThroughputBps)
	pusher.Collector(uplinkThroughputBps)
	pusher.Collector(snr)
	pusher.Collector(popPingLatencyMs)
	pusher.Collector(popPingDropRate)
	pusher.Collector(percentObstructed)
	pusher.Collector(secondsObstructed)

	if err := pusher.Add(); err != nil {
		logger.Error(err)
		storeDataResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, storeDataResponse)
		return
	}

	if err := stores.GetStore().SetDishyLastReceivedStats(context.Background(), dishy.ID, when); err != nil {
		logger.Error(err)
		storeDataResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, storeDataResponse)
		return
	}

	fmt.Println("data sent to push gateway")
	w.WriteHeader(http.StatusCreated)
}
