package auth

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

//TODO revoir la godoc
//TODO SPE revoir la visibilite de la plupart des fields et methods

//BasicAuth represents a HTTP Request with Basic Auth
type BasicAuth struct {
	Username string
	Password string
}

//Get invokes a HTTP Get with BasicAuth and handles errors
func (r BasicAuth) Get(url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(r.Username, r.Password)

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
