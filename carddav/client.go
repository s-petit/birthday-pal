package carddav

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Contacts calls a CardDAV server with an URL and BasicAuth
func Contacts(url string, username string, password string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		return "", errors.New("a unexpected error occurred during connexion to CardDAV server - http code is " + strconv.Itoa(resp.StatusCode))
	}

	return string(body), nil
}
