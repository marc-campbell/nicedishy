package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func StreamDishies(w http.ResponseWriter, r *http.Request) {
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
		for {
			messageChan <- "data: Hello, world!\n\n"
			time.Sleep(time.Second)
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
