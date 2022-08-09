package handlers

import (
	"net/http"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	tokentypes "github.com/marc-campbell/nicedishy/pkg/token/types"
)

func RequireValidInternalAuthQuietMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, err := requireValidInternalAuth(w, r)
		if err != nil {
			return
		}

		r = setInternalUser(r, username)
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
