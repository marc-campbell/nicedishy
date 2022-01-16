package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marc-campbell/nicedishy/pkg/logger"
)

type UpdateWaitlistRequest struct {
	EmailAddress string `json:"emailAddress"`
}

type UpdateWaitlistResponse struct {
	Error string `json:"error,omitempty"`
}

func UpdateWaitlist(w http.ResponseWriter, r *http.Request) {
	response := UpdateWaitlistResponse{}

	request := UpdateWaitlistRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	JSON(w, http.StatusOK, response)
}
