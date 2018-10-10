package auth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const tokenFile = "token.json"
const configFile = "config.json"

//OAuth2Authenticator is used to perform OAuth2Authenticator authentication
type OAuth2Authenticator struct {
	Scope   string
	Profile OAuthProfile
}

//Client returns a HTTP client authenticated with OAuth2Authenticator
func (oa OAuth2Authenticator) Client() (*http.Client, error) {

	// config returns the configuration from client_secret.json
	config, err := oa.config()
	if err != nil {
		return nil, err
	}
	// load the token from the cache
	token, err := oa.Profile.loadProfileTokenFromCache()
	if err != nil {
		fmt.Println("not authenticated yet! please use 'birthday-pal oauth' command.")
		return nil, err
	}

	// Create the API client with a background context.
	ctx := context.Background()
	client := config.Client(ctx, token)

	return client, nil
}

//Authenticate parses and save the config file provided by user
//then performs a an OAuth2Authenticator authentication
//then saves token in cache
func (oa OAuth2Authenticator) Authenticate(configFilePath string) error {

	oa.Profile.saveProfileConfigInCache(configFilePath)

	config, err := oa.config()
	if err != nil {
		return err
	}

	// Compute the URL for the authorization
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	// Notify the user of the web browser.
	fmt.Println("Please use a browser in order to let birthday-pal access your contacts.")

	err = oa.Profile.System.OpenBrowser(authURL)

	// If we couldn't open the web browser, prompt the user to do it manually.
	if err != nil {
		fmt.Printf("Copy and paste the following link: \n%s\n\n", authURL)
	}

	// prompt for the authorization code
	code, err := oa.Profile.System.Prompt()
	if err != nil {
		return fmt.Errorf("unable to read authorization code %v", err)
	}

	// Perform the exchange for the token
	token, err := oa.Profile.System.ExchangeToken(config, code)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web %v", err)
	}

	// Cache the config and the token to disk
	oa.Profile.saveProfileTokenInCache(token)

	return nil
}

// config loads the oauth config file from the cache. It is used both to
// create the client for requests as well as to perform authentication.
func (oa OAuth2Authenticator) config() (*oauth2.Config, error) {

	configFile := filepath.Join(oa.Profile.profileCachePath(), configFile)

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(data, oa.Scope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	return config, nil
}
