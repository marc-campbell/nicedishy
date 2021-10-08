package handlers

import (
	"net/http"
)

type HealthzResponse struct {
	Version string         `json:"version"`
	GitSHA  string         `json:"gitSha"`
	Status  StatusResponse `json:"status"`
}

type StatusResponse struct {
}

// Healthz route is UNAUTHENTICATED
func Healthz(w http.ResponseWriter, r *http.Request) {

	statusCode := 200
	healthzResponse := HealthzResponse{}
	JSON(w, statusCode, healthzResponse)
}
