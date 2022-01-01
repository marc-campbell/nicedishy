package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	nats "github.com/nats-io/nats.go"
)

type StreamPublicStatsResponse struct {
	ConnectedDishyCount  int     `json:"connectedDishyCount"`
	AllTimeDishyCount    int     `json:"allTimeDishyCount"`
	NewDishyCount        int     `json:"newDishyCount"`
	HighestDownloadSpeed float64 `json:"highestDownloadSpeed"`
	AverageDownloadSpeed float64 `json:"averageDownloadSpeed"`
	LowestPingTime       float64 `json:"lowestPingTime"`
	AveragePingTime      float64 `json:"averagePingTime"`
}

func StreamPublicStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ALLOW_ORIGIN"))

	messageChan := make(chan string)
	flusher := w.(http.Flusher)

	defer func() {
		close(messageChan)
		messageChan = nil
	}()

	go func() {
		opts := []nats.Option{nats.Name("public stats")}
		totalWait := 10 * time.Minute
		reconnectDelay := time.Second

		opts = append(opts, nats.ReconnectWait(reconnectDelay))
		opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
		opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
		}))
		opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("Reconnected [%s]", nc.ConnectedUrl())
		}))
		opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
			log.Fatalf("Exiting: %v", nc.LastError())
		}))

		nc, err := nats.Connect(os.Getenv("NATS_ENDPOINT"), opts...)
		if err != nil {
			log.Fatal(err)
		}

		subj := "public_stats"

		nc.Subscribe(subj, func(msg *nats.Msg) {
			messageChan <- string(msg.Data)
		})
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		}

		runtime.Goexit()
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
