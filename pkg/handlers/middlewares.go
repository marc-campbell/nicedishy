package handlers

import (
	"net/http"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	tokentypes "github.com/marc-campbell/nicedishy/pkg/token/types"
)

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
