package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/session"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	"github.com/marc-campbell/nicedishy/pkg/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type LoginCallbackRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type LoginCallbackResponse struct {
	Error       string `json:"error,omitempty"`
	Token       string `json:"token,omitempty"`
	RedirectURI string `json:"redirectUri,omitempty"`
}

func LoginCallback(w http.ResponseWriter, r *http.Request) {
	loginCallbackResponse := LoginCallbackResponse{}

	loginCallbackRequest := LoginCallbackRequest{}
	if err := json.NewDecoder(r.Body).Decode(&loginCallbackRequest); err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}

	ok, next, err := stores.GetStore().GetOAuthState(context.Background(), loginCallbackRequest.State)
	if err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}

	if !ok {
		loginCallbackResponse.Error = "Invalid state"
		JSON(w, http.StatusBadRequest, loginCallbackResponse)
		return
	}

	if next == "undefined" {
		next = ""
	}
	
	loginCallbackResponse.RedirectURI = next

	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENTID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENTSECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECTURI"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	tok, err := conf.Exchange(oauth2.NoContext, loginCallbackRequest.Code)
	if err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}
	client := conf.Client(oauth2.NoContext, tok)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}

	userInfo := map[string]interface{}{}
	if err := json.Unmarshal(data, &userInfo); err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}

	emailAddress := userInfo["email"].(string)
	avatar := userInfo["picture"].(string)

	// get or create the local user account
	user, err := user.GetOrCreate(context.TODO(), emailAddress, avatar)
	if err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}

	sess, err := session.CreateSessionForUser(context.TODO(), user, tok.AccessToken)
	if err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}

	token, err := session.Token(context.TODO(), sess)
	if err != nil {
		logger.Error(err)
		loginCallbackResponse.Error = err.Error()
		JSON(w, http.StatusInternalServerError, loginCallbackResponse)
		return
	}

	loginCallbackResponse.Token = token
	JSON(w, http.StatusOK, loginCallbackResponse)
	return

}
