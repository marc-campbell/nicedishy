package handlers

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type LoginResponse struct {
	RedirectURL string `json:"redirectURL"`
	Error       string `json:"error"`
}

// Login is the handler called from the client when they click the login button
// This should generate the GitHub URL that we want to redirect the client to
// and return it
func Login(w http.ResponseWriter, r *http.Request) {
	loginResponse := LoginResponse{}

	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENTID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENTSECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECTURI"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	url := conf.AuthCodeURL("TODO")

	loginResponse.RedirectURL = url

	JSON(w, http.StatusOK, loginResponse)
}
