package handlers

import (
	"context"
	"net/http"

	"github.com/marc-campbell/nicedishy/pkg/dishy"
)

type DeployLatestPublicDashboardsResponse struct {
	Error string `json:"error,omitempty"`
}

func DeployLatestPublicDashboards(w http.ResponseWriter, r *http.Request) {

	response := DeployLatestPublicDashboardsResponse{}

	if err := dishy.UpdatePublicDashboards(context.TODO()); err != nil {
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	JSON(w, http.StatusOK, response)
}
