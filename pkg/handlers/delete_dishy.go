package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

type DeleteDishyResponse struct {
	Error string `json:"error,omitempty"`
}

func DeleteDishy(w http.ResponseWriter, r *http.Request) {
	deleteDishyResponse := DeleteDishyResponse{}

	userID := getUserID(r)
	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	dishy, err := stores.GetStore().GetDishyForUser(context.TODO(), mux.Vars(r)["id"], userID)
	if err != nil {
		logger.Error(err)
		deleteDishyResponse.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if dishy == nil {
		deleteDishyResponse.Error = "dishy not found"
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := stores.GetStore().DeleteDishy(context.TODO(), mux.Vars(r)["id"]); err != nil {
		logger.Error(err)
		deleteDishyResponse.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO delete the stats÷

	w.WriteHeader(http.StatusNoContent)
}
