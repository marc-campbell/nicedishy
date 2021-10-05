package handlers

import (
	"context"
	"net/http"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

type GetDishiesResponse struct {
	Dishies []*dishytypes.Dishy `json:"dishies"`
	Error   string              `json:"error,omitempty"`
}

func GetDishies(w http.ResponseWriter, r *http.Request) {
	getDishiesResponse := GetDishiesResponse{}

	userID := getUserID(r)
	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	dishies, err := stores.GetStore().ListDishies(context.TODO(), userID)
	if err != nil {
		logger.Error(err)
		getDishiesResponse.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	getDishiesResponse.Dishies = dishies
	JSON(w, http.StatusOK, getDishiesResponse)
}
