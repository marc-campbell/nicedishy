package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
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

	userID := getUserID(r)
	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	dishy, err := stores.GetStore().CreateDishy(context.TODO(), userID, createDishyRequest.Name)
	if err != nil {
		logger.Error(err)
		createDishyResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, createDishyResponse)
		return
	}

	createDishyResponse.Dishy = dishy
	JSON(w, http.StatusCreated, createDishyResponse)
}
