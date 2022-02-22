package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marc-campbell/nicedishy/pkg/dishy"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

type DeployLatestDishyDashboardResponse struct {
	Error string `json:"error,omitempty"`
}

func DeployLatestDishyDashboard(w http.ResponseWriter, r *http.Request) {
	dishyID := mux.Vars(r)["id"]

	response := DeployLatestDishyDashboardResponse{}

	d, err := stores.GetStore().GetDishy(context.TODO(), dishyID)
	if err != nil {
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	if err := dishy.UpdateGrafanaDashboard(context.TODO(), d.ID, d.Name); err != nil {
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	JSON(w, http.StatusOK, response)
}
