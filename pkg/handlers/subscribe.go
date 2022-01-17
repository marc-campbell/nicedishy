package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/mailer"
)

type SubscribeRequest struct {
	EmailAddress string `json:"emailAddress"`
}

type SubscribeResponse struct {
	Error string `json:"error,omitempty"`
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	response := SubscribeResponse{}

	request := SubscribeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	if request.EmailAddress == "" {
		response.Error = "Invalid email address"
		JSON(w, http.StatusBadRequest, response)
		return
	}

	err := mailer.SignUpForNewsletter(r.Context(), request.EmailAddress, true)
	if err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	JSON(w, http.StatusOK, response)
}
