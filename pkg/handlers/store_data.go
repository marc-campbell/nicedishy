package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
	"github.com/marc-campbell/nicedishy/pkg/stores"
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

	metricsDB := persistence.MustGetMetricsDBSession()
	query := `insert into dishy_data (time, dishy_id, state, snr, downlink_throughput_bps, uplink_throughput_bps, pop_ping_latency_ms, pop_ping_drop_rate, percent_obstructed, seconds_obstructed) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err = metricsDB.Exec(context.Background(), query, when, dishy.ID, storeDataRequest.Status.State, storeDataRequest.Status.SNR, storeDataRequest.Status.DownlinkThroughputBps, storeDataRequest.Status.UplinkThroughputBps, storeDataRequest.Status.PopPingLatencyMs, storeDataRequest.Status.PopPingDropRate, storeDataRequest.Status.PercentObstructed, storeDataRequest.Status.SecondsObstructed)
	if err != nil {
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

	w.WriteHeader(http.StatusCreated)
}
