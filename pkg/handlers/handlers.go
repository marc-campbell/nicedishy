package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/marc-campbell/nicedishy/pkg/logger"
)

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type, origin, accept, authorization")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func StreamJSON(c *websocket.Conn, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Error(err)
		return
	}

	err = c.WriteMessage(websocket.TextMessage, response)
	if err != nil {
		logger.Error(err)
		return
	}
}
