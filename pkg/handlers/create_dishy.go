package handlers

import (
	"encoding/json"
	"net/http"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/logger"
)

type CreateDishyRequest struct {
	Name string `json:"name"`
}

type CreateDishyResponse struct {
	Error string            `json:"error,omitempty"`
	Dishy *dishytypes.Dishy `json:"dishy,omitempty"`
}

func CreateDishy(w http.ResponseWriter, r *http.Request) {
	createDishyResponse := CreateDishyResponse{}

	createDishyRequest := CreateDishyRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createDishyRequest); err != nil {
		logger.Error(err)
		createDishyResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, createDishyResponse)
		return
	}

	createDishyResponse.Error = "Not implemented"
	JSON(w, http.StatusNotImplemented, createDishyResponse)
}
