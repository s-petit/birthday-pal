package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/s-petit/birthday-pal/system"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const tokenFile = "token.json"
const configFile = "config.json"

//OAuth2 is used to perform OAuth2 authentication
type OAuth2 struct {
	Scope   string
	System  system.System
	Profile string
}

//Client returns a HTTP client authenticated with OAuth2
func (oa OAuth2) Client() (*http.Client, error) {

	// config returns the configuration from client_secret.json
	config, err := oa.config()
	if err != nil {
		return nil, err
	}
	// load the token from the cache
	token, err := oa.loadTokenFromCache()
	if err != nil {
		fmt.Println("not authenticated yet! please use 'birthday-pal oauth' command.")
		return nil, err
	}

	// Create the API client with a background context.
	ctx := context.Background()
	client := config.Client(ctx, token)

	return client, nil
}

// save the token to the specified path.
func (oa OAuth2) saveTokenInCache(token *oauth2.Token) error {

	tokenPath := filepath.Join(oa.System.CachePath(oa.Profile), tokenFile)

	// Open the file for writing
	tokenFile, err := os.Create(tokenPath)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer tokenFile.Close()

	// Encode the token and write to disk
	if err := json.NewEncoder(tokenFile).Encode(token); err != nil {
		return fmt.Errorf("could not encode oauth token: %v", err)
	}

	return nil
}

// save the token to the specified path.
func (oa OAuth2) saveConfigInCache(secretPath string) error {

	data, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %v", err)
	}

	// TODO SPE oa.oa ?? peut mieux faire ?
	conf := filepath.Join(oa.System.CachePath(oa.Profile), configFile)
	ioutil.WriteFile(conf, data, 0644)

	return nil
}

// load the oauth2 token from the specified path
func (oa OAuth2) loadTokenFromCache() (*oauth2.Token, error) {

	cachePath := oa.System.CachePath(oa.Profile)
	tokenPath := filepath.Join(cachePath, tokenFile)

	// Open the file at the path
	f, err := os.Open(tokenPath)
	if err != nil {
		return nil, fmt.Errorf("could not open cache file at %s: %v", tokenPath, err)
	}
	defer f.Close()

	// Decode the JSON token cache
	token := new(oauth2.Token)
	if err := json.NewDecoder(f).Decode(token); err != nil {
		return nil, fmt.Errorf("could not decode token in cache file at %s: %v", tokenPath, err)
	}

	return token, nil
}

//Authenticate performs a an OAuth2 authentication then save config and token in cache
func (oa OAuth2) Authenticate(secretPath string) error {

	oa.saveConfigInCache(secretPath)

	config, err := oa.config()
	if err != nil {
		return err
	}

	// Compute the URL for the authorization
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	// Notify the user of the web browser.
	fmt.Println("Please use a browser in order to let birthday-pal access your contacts.")

	err = oa.System.OpenBrowser(authURL)

	// If we couldn't open the web browser, prompt the user to do it manually.
	if err != nil {
		fmt.Printf("Copy and paste the following link: \n%s\n\n", authURL)
	}

	// prompt for the authorization code
	code, err := oa.System.Prompt()
	if err != nil {
		return fmt.Errorf("unable to read authorization code %v", err)
	}

	// Perform the exchange for the token
	token, err := oa.System.ExchangeToken(config, code)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web %v", err)
	}

	// Cache the config and the token to disk
	oa.saveTokenInCache(token)

	return nil
}

//TODO SPE faire attention aux commentaires de cette classe
// et aux noms a utiliser ? token ? secret ? config ?

// config loads the secret file from the SecretPath. It is used both to
// create the client for requests as well as to perform authentication.
func (oa OAuth2) config() (*oauth2.Config, error) {

	configFile := filepath.Join(oa.System.CachePath(oa.Profile), configFile)

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
