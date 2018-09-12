package auth

import (
	"context"
	"net/http"
)

//TODO revoir la godoc

//OAuth2 represents a HTTP Request with OAuth2
type OAuth2 struct {
	Scope      string
	SecretPath string
}

func (oa OAuth2) Authenticate() error {
	return oa.authentication().authenticate()
}

func (oa OAuth2) Client() (*http.Client, error) {
	// Initialize authentication
	auth := oa.authentication()

	// load the configuration from client_secret.json
	config, err := auth.config()
	if err != nil {
		return nil, err
	}
	// load the token from the cache or force authentication
	token, err := auth.getToken()
	if err != nil {
		return nil, err
	}

	// Create the API client with a background context.
	ctx := context.Background()
	client := config.Client(ctx, token)

	return client, nil
}

//Get invokes a HTTP Get with BasicAuth and handles errors
/*func (oa OAuth2) Call(url string) (string, error) {

	client, err := oa.Client()
	if err != nil {
		return "", err
	}

	resp, err := client.Get(url)

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
}*/

func (oa OAuth2) authentication() *authentication {
	return &authentication{Scope: oa.Scope, SecretPath: oa.SecretPath}
}
