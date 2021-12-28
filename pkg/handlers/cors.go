package handlers

import (
	"net/http"
	"os"
)

func CORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ALLOW_ORIGIN"))
	w.Header().Set("Access-Control-Allow-Headers", "content-type, origin, accept, authorization")
}
