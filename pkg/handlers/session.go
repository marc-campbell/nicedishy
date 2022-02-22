package handlers

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/marc-campbell/nicedishy/pkg/session"
	sessiontypes "github.com/marc-campbell/nicedishy/pkg/session/types"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	tokentypes "github.com/marc-campbell/nicedishy/pkg/token/types"
	"github.com/pkg/errors"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type tokenKey struct{}
type sessionKey struct{}
type internalUsernameKey struct{}

// getUserID will return the user id, regardless if it's a token or a session
func getUserID(r *http.Request) string {
	if getSession(r) != nil {
		return getSession(r).UserID
	}

	return ""
}

func setInternalUser(r *http.Request, username string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), internalUsernameKey{}, username))
}

func setSession(r *http.Request, sess *sessiontypes.Session) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), sessionKey{}, sess))
}

func getSession(r *http.Request) *sessiontypes.Session {
	val := r.Context().Value(sessionKey{})
	sess, _ := val.(*sessiontypes.Session)
	return sess
}

func setToken(r *http.Request, tok *tokentypes.Token) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), tokenKey{}, tok))
}

func getToken(r *http.Request) *tokentypes.Token {
	val := r.Context().Value(tokenKey{})
	tok, _ := val.(*tokentypes.Token)
	return tok
}

func optionalValidToken(w http.ResponseWriter, r *http.Request) (*tokentypes.Token, error) {
	auth := r.Header.Get("authorization")

	if auth == "" {
		return nil, errors.New("authorization header empty")
	}

	parts := strings.Fields(auth)
	if len(parts) != 2 {
		return nil, errors.New("expected 2 parts in auth header")
	}

	if parts[0] != "Token" {
		return nil, errors.New("not a token auth")
	}

	t, err := stores.GetStore().GetToken(context.Background(), parts[1])
	if err != nil {
		return nil, errors.New("no token found")
	}

	return t, nil
}

func requireValidInternalAuth(w http.ResponseWriter, r *http.Request) (string, error) {
	auth := r.Header.Get("authorization")

	if auth == "" {
		return "", errors.New("authorization header empty")
	}

	parts := strings.Fields(auth)
	if len(parts) != 2 {
		return "", errors.New("expected 2 parts in auth header")
	}

	if parts[0] != "Basic" {
		return "", errors.New("not a bearer auth")
	}

	decoded, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", errors.New("invalid base64")
	}

	usernameAndPassword := strings.Split(string(decoded), ":")
	if len(usernameAndPassword) != 2 {
		return "", errors.New("invalid username and password")
	}

	username := usernameAndPassword[0]
	password := usernameAndPassword[1]

	if username == "retool" && password == os.Getenv("NICEDISHY_RETOOL_API_PASSWORD") {
		return username, nil
	}

	return "", errors.New("invalid auth")
}

func requireValidToken(w http.ResponseWriter, r *http.Request) (*tokentypes.Token, error) {
	tok, err := optionalValidToken(w, r)
	if err != nil {
		response := ErrorResponse{Error: err.Error()}
		JSON(w, http.StatusUnauthorized, response)
		return nil, err
	}

	return tok, nil
}

func optionalValidSession(w http.ResponseWriter, r *http.Request) (*sessiontypes.Session, error) {
	auth := r.Header.Get("authorization")

	if auth == "" {
		return nil, errors.New("authorization header empty")
	}

	parts := strings.Fields(auth)
	if len(parts) != 2 {
		return nil, errors.New("expected 2 parts in auth header")
	}

	if parts[0] != "Bearer" {
		return nil, errors.New("not a bearer auth")
	}

	s, err := session.Parse(auth)
	if err != nil {
		return nil, errors.New("no session found")
	}

	return s, nil
}

func requireValidSession(w http.ResponseWriter, r *http.Request) (*sessiontypes.Session, error) {
	sess, err := optionalValidSession(w, r)
	if err != nil {
		response := ErrorResponse{Error: err.Error()}
		JSON(w, http.StatusUnauthorized, response)
		return nil, err
	}

	return sess, err
}
