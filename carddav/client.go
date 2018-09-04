package carddav

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Request represents a HTTP Request
type Request interface {
	Get() (string, error)
}

//BasicAuthRequest represents a HTTP Request with Basic Auth
type BasicAuthRequest struct {
	URL      string
	Username string
	Password string
}

func (c BasicAuthRequest) request() (*http.Request, error) {
	req, err := http.NewRequest("GET", c.URL, nil)
	req.SetBasicAuth(c.Username, c.Password)
	return req, err
}

//Get invokes a HTTP Get with BasicAuth and handles errors
func (c BasicAuthRequest) Get() (string, error) {

	req, err := c.request()
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
