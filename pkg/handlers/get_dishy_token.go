package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

type GetDishyTokenResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

func GetDishyToken(w http.ResponseWriter, r *http.Request) {
	getDishyTokenResponse := GetDishyTokenResponse{}

	userID := getUserID(r)
	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	dishy, err := stores.GetStore().GetDishyForUser(context.TODO(), mux.Vars(r)["id"], userID)
	if err != nil {
		logger.Error(err)
		getDishyTokenResponse.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if dishy == nil {
		getDishyTokenResponse.Error = "dishy not found"
		w.WriteHeader(http.StatusNotFound)
		return
	}

	token, err := stores.GetStore().CreateDishyToken(context.TODO(), dishy.ID)
	if err != nil {
		logger.Error(err)
		getDishyTokenResponse.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	getDishyTokenResponse.Token = token
	JSON(w, http.StatusCreated, getDishyTokenResponse)
}
