package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/dishy"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	"go.uber.org/zap"
)

type StoreSpeedSpeedRequest struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
}

type StoreSpeedRequest struct {
	When  string                 `json:"when"`
	Speed StoreSpeedSpeedRequest `json:"speed"`
}

type StoreSpeedResponse struct {
	Error string `json:"error,omitempty"`
}

func StoreSpeed(w http.ResponseWriter, r *http.Request) {
	response := StoreSpeedResponse{}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		JSON(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("UA %s payload: %s\n", r.Header.Get("user-agent"), payload)

	storeSpeedRequest := StoreSpeedRequest{}
	if err := json.Unmarshal(payload, &storeSpeedRequest); err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	ipAddress := r.Header.Get("True-Client-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			logger.Error(err)
			response.Error = err.Error()
			JSON(w, http.StatusInternalServerError, response)
			return
		}

		ipAddress = ip
	}

	// 206.214.226.67,172.70.210.63
	// split and takje the first one
	if strings.Contains(ipAddress, ",") {
		ipAddress = strings.Split(ipAddress, ",")[0]
	}

	// update the "last received data from" date of the dish
	d := DishyFromTokenContext(r)
	if d == nil {
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	when, err := time.Parse(time.RFC3339, storeSpeedRequest.When)
	if err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	// if it's been over 30 days, lets do a new geocheck of this source
	if d.LastGeocheckAt == nil || time.Since(*d.LastGeocheckAt) > 30*24*time.Hour {
		logger.Info("geochecking",
			zap.String("dishyID", d.ID))

		geocheck, err := dishy.Geocheck(d.ID, ipAddress)
		if err != nil {
			logger.Error(err)
			response.Error = err.Error()
			JSON(w, http.StatusInternalServerError, response)
			return
		}

		if geocheck.Org != "SpaceX Services, Inc." {
			fmt.Printf("THIS IS NOT A DISHY: org = %s\n", geocheck.Org)
		}

		if err := stores.GetStore().UpdateDishyGeo(context.TODO(), d.ID, when, geocheck); err != nil {
			logger.Error(err)
			response.Error = err.Error()
			JSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	metricsDB := persistence.MustGetMetricsDBSession()
	query := `insert into dishy_data (
time, dishy_id, ip_address,
download_speed, upload_speed)
values
($1, $2, $3, $4, $5)`
	_, err = metricsDB.Exec(context.Background(), query, when, d.ID, ipAddress,
		storeSpeedRequest.Speed.Download, &storeSpeedRequest.Speed.Upload)
	if err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	if err := stores.GetStore().SetDishyLastReceivedStats(context.Background(), d.ID, when); err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
