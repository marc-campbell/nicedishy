package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/marc-campbell/nicedishy/pkg/logger"
)

type SubscribeRequest struct {
	EmailAddress string `json:"emailAddress"`
}

type SubscribeResponse struct {
	Error string `json:"error,omitempty"`
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	response := SubscribeResponse{}

	request := SubscribeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	if request.EmailAddress == "" {
		response.Error = "Invalid email address"
		JSON(w, http.StatusBadRequest, response)
		return
	}

	body := map[string]string{
		"Email": request.EmailAddress,
	}
	b, err := json.Marshal(body)
	if err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	// subscribe to the newsletter in mailchimp
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.moosend.com/v3/subscribers/%s/subscribe.json?apikey=%s", os.Getenv("MOOSEND_LISTID"), os.Getenv("MOOSEND_API_KEY")), ioutil.NopCloser(bytes.NewReader(b)))
	if err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	if resp.StatusCode != http.StatusOK {
		response.Error = fmt.Sprintf("%d: %s", resp.StatusCode, resp.Status)
		JSON(w, resp.StatusCode, response)
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	responseDecoded := map[string]interface{}{}
	if err := json.Unmarshal(responseBody, &responseDecoded); err != nil {
		logger.Error(err)
		response.Error = err.Error()
		JSON(w, http.StatusInternalServerError, response)
		return
	}

	JSON(w, http.StatusOK, response)
}
