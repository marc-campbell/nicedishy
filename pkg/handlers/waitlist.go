package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

type UpdateWaitlistRequest struct {
	AlreadyHaveDishy string   `json:"alreadyHaveDishy"`
	HowLongWithDishy string   `json:"howLongWithDishy"`
	PrimaryOrBackup  string   `json:"primaryOrBackup"`
	OperatingSystems []string `json:"operatingSystems"`
	WhyAccess        []string `json:"whyAccess"`
}

type UpdateWaitlistResponse struct {
	Error string `json:"error,omitempty"`
}

func UpdateWaitlist(w http.ResponseWriter, r *http.Request) {
	response := UpdateWaitlistResponse{}

	userID := getUserID(r)
	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	request := UpdateWaitlistRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	if err := stores.GetStore().UpdateUserWaitlistData(context.TODO(), userID, request.AlreadyHaveDishy, request.HowLongWithDishy, request.PrimaryOrBackup, request.OperatingSystems, request.WhyAccess); err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	JSON(w, http.StatusOK, response)
}
