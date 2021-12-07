package handlers

import (
	"context"
	"net/http"

	"github.com/marc-campbell/nicedishy/pkg/dishy"
	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

type GetDishiesResponse struct {
	Dishies []*dishytypes.DishyWithStats `json:"dishies"`
	Error   string                       `json:"error,omitempty"`
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

	// attach some stats
	dishiesWithStats := []*dishytypes.DishyWithStats{}
	for _, d := range dishies {
		dishyWithStats := dishytypes.DishyWithStats{
			Dishy: *d,
		}

		latestStats, err := dishy.GetLatestStats(d.ID)
		if err != nil {
			logger.Error(err)
			getDishiesResponse.Error = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		dishyWithStats.Latest = latestStats

		recentStats, err := dishy.GetRecentStats(d.ID)
		if err != nil {
			logger.Error(err)
			getDishiesResponse.Error = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		dishyWithStats.RecentStats = recentStats

		dishiesWithStats = append(dishiesWithStats, &dishyWithStats)
	}

	getDishiesResponse.Dishies = dishiesWithStats
	JSON(w, http.StatusOK, getDishiesResponse)
}
