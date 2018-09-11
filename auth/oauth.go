package auth

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//TODO revoir la godoc

//OAuth2 represents a HTTP Request with OAuth2
type OAuth2 struct {
	Scope      string
	SecretPath string
}

func (oa OAuth2) authentication() *authentication {
	return &authentication{Scope: oa.Scope, SecretPath: oa.SecretPath}
}

func (oa OAuth2) Authenticate() error {
	return oa.authentication().authenticate()
}

func (oa OAuth2) Clt(url string) (*http.Client, error) {
	return oa.oauthClient(), nil
}

func (oa OAuth2) oauthClient() *http.Client {
	// Initialize authentication
	auth := oa.authentication()

	// load the configuration from client_secret.json
	config, err := auth.config()
	if err != nil {
		log.Fatal(err.Error())
	}

	// load the token from the cache or force authentication
	token, err := auth.getToken()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create the API client with a background context.
	ctx := context.Background()
	client := config.Client(ctx, token)

	return client
}

//Get invokes a HTTP Get with BasicAuth and handles errors
func (oa OAuth2) Get(url string) (string, error) {

	resp, err := oa.oauthClient().Get(url)

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

//Get invokes a HTTP Get with BasicAuth and handles errors
func (oa OAuth2) GetRes(url string) (*http.Response, error) {
	return oa.oauthClient().Get(url)
}
