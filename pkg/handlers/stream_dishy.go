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

type StreamDishyResponse struct {
	Dishy    *dishytypes.Dishy                    `json:"dishy"`
	Stats    map[time.Time]*dishytypes.DishyStat  `json:"stats"`
	Speeds   map[time.Time]*dishytypes.DishySpeed `json:"speeds"`
	Versions map[string]string                    `json:"versions"`
}

func StreamDishy(w http.ResponseWriter, r *http.Request) {
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
			d, err := stores.GetStore().GetDishyForUser(context.TODO(), r.URL.Query().Get("id"), userID)
			if err != nil {
				logger.Error(err)
				continue
			}

			stats, speeds, err := dishy.GetRecentStats(d.ID)
			if err != nil {
				logger.Error(err)
				continue
			}

			softwareVersion, hardwareVersion, err := stores.GetStore().GetDishyVersions(context.TODO(), d.ID)
			if err != nil {
				logger.Error(err)
				continue
			}

			message := StreamDishyResponse{
				Dishy:  d,
				Stats:  stats,
				Speeds: speeds,
				Versions: map[string]string{
					"hardware": hardwareVersion,
					"software": softwareVersion,
				},
			}

			b, err := json.Marshal(message)
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
