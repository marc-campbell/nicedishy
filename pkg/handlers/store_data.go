package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marc-campbell/nicedishy/pkg/logger"
)

type StoreDataStatusDeviceInfoRequest struct {
	HardwareVersion string `json:"hardwareVersion"`
	SoftwareVersion string `json:"softwareVersion"`
}

type StoreDataStatusRequest struct {
	DeviceInfo StoreDataStatusDeviceInfoRequest `json:"deviceInfo"`
}

type StoreDataRequest struct {
	When   string                 `json:"when"`
	Status StoreDataStatusRequest `json:"status"`
}

type StoreDataResponse struct {
	Error       string `json:"error,omitempty"`
	Token       string `json:"token,omitempty"`
	RedirectURI string `json:"redirectUri,omitempty"`
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

	fmt.Printf("%#v\n", storeDataRequest)

	w.WriteHeader(http.StatusCreated)
}
