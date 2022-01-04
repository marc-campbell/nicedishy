package handlers

import (
	"context"
	"net/http"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	usertypes "github.com/marc-campbell/nicedishy/pkg/user/types"
)

type WhoAmIResponse struct {
	Error string          `json:"error,omitempty"`
	User  *usertypes.User `json:"user,omitempty"`
}

func WhoAmI(w http.ResponseWriter, r *http.Request) {
	sess := getSession(r)
	u, err := stores.GetStore().GetUserByID(context.TODO(), sess.UserID)
	if err != nil {
		logger.Error(err)
		JSON(w, http.StatusInternalServerError, WhoAmIResponse{Error: err.Error()})
		return
	}

	response := WhoAmIResponse{
		User: u,
	}

	JSON(w, http.StatusCreated, response)
}
