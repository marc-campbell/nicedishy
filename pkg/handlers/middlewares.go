package handlers

import (
	"context"
	"net/http"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	tokentypes "github.com/marc-campbell/nicedishy/pkg/token/types"
)

func RequireValidNonceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonceID := r.URL.Query().Get("nonce")
		sessionID, err := stores.GetStore().GetSessionNonce(context.TODO(), nonceID)
		if err != nil {
			return
		}

		sess, err := stores.GetStore().GetSession(context.TODO(), sessionID)
		if err != nil {
			return
		}

		r = setSession(r, sess)
		next.ServeHTTP(w, r)
	})
}

func RequireValidSessionQuietMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, err := requireValidSession(w, r)
		if err != nil {
			return
		}

		r = setSession(r, s)
		next.ServeHTTP(w, r)
	})
}

func RequireValidTokenQuietMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := requireValidToken(w, r)
		if err != nil {
			return
		}

		r = setToken(r, t)
		next.ServeHTTP(w, r)
	})
}

func DishyFromTokenContext(r *http.Request) *dishytypes.Dishy {
	val := r.Context().Value(tokenKey{})
	tok, ok := val.(*tokentypes.Token)
	if !ok {
		return nil
	}
	return tok.Dishy
}
