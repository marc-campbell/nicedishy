package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/dishy"
	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

func StreamDishies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ALLOW_ORIGIN"))

	userID := getUserID(r)
	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	messageChan := make(chan string)
	flusher := w.(http.Flusher)

	defer func() {
		close(messageChan)
		messageChan = nil
	}()

	go func() {
		for {
			dishies, err := stores.GetStore().ListDishies(context.TODO(), userID)
			if err != nil {
				logger.Error(err)
				continue
			}

			fmt.Printf("dishies: %+v\n", dishies)

			// attach some stats and speeds
			dishiesWithStats := []*dishytypes.DishyWithStats{}
			for _, d := range dishies {
				dishyWithStats := dishytypes.DishyWithStats{
					Dishy: *d,
				}

				latestStats, err := dishy.GetLatestStats(d.ID)
				if err != nil {
					logger.Error(err)
					continue
				}

				fmt.Printf("latest stats: %+v\n", latestStats)
				dishyWithStats.LatestStats = latestStats

				latestSpeeds, err := dishy.GetLatestSpeeds(d.ID)
				if err != nil {
					logger.Error(err)
					continue
				}
				fmt.Printf("latest speeds: %+v\n", latestSpeeds)
				dishyWithStats.LatestSpeeds = latestSpeeds

				dishiesWithStats = append(dishiesWithStats, &dishyWithStats)
			}

			b, err := json.Marshal(dishiesWithStats)
			if err != nil {
				logger.Error(err)
				continue
			}

			messageChan <- string(b)

			time.Sleep(time.Second * 10)
		}
	}()

	for {
		select {
		case message := <-messageChan:
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()

		case <-r.Context().Done():
			return
		}
	}

}
