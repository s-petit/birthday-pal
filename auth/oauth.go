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

//Authenticate performs a an OAuth2 authentication
func (oa OAuth2) Authenticate() error {
	return oa.authentication().authenticate()
}

//Client returns a HTTP client authenticated with OAuth2
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

func (oa OAuth2) authentication() *authentication {
	return &authentication{Scope: oa.Scope, SecretPath: oa.SecretPath}
}
