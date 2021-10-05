package handlers

import (
	"net/http"
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
