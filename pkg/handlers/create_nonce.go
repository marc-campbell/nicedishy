package handlers

import (
	"context"
	"net/http"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

type CreateNonceResponse struct {
	Error string `json:"error,omitempty"`
	Nonce string `json:"nonce,omitempty"`
}

func CreateNonce(w http.ResponseWriter, r *http.Request) {
	sess := getSession(r)
	nonce, err := stores.GetStore().CreateSessionNonce(context.TODO(), sess.ID)
	if err != nil {
		logger.Error(err)
		JSON(w, http.StatusInternalServerError, CreateNonceResponse{Error: err.Error()})
		return
	}

	response := CreateNonceResponse{
		Nonce: nonce,
	}

	JSON(w, http.StatusCreated, response)
}
