package carddav

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Request holds methods necessary for requesting cardDAV HTTP servers.
type Request interface {
	Get() (string, error)
}

//BasicAuthRequest represents a CardDAV HTTP Request with Basic Auth
type BasicAuthRequest struct {
	URL      string
	Username string
	Password string
}

func (r BasicAuthRequest) request() (*http.Request, error) {
	httpRequest, err := http.NewRequest("GET", r.URL, nil)
	httpRequest.SetBasicAuth(r.Username, r.Password)
	return httpRequest, err
}

//Get invokes a HTTP Get with BasicAuth and handles errors
func (r BasicAuthRequest) Get() (string, error) {

	req, err := r.request()
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("a unexpected error occurred during connexion to CardDAV server - http code is " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
