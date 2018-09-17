package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/s-petit/birthday-pal/system"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

//OAuth2 represents a HTTP Request with OAuth2
type OAuth2 struct {
	Scope      string
	SecretPath string
	//TODO revoir ce champ et mocker
	//cachePath string
	system system.System
}

//Client returns a HTTP client authenticated with OAuth2
func (oa OAuth2) Client() (*http.Client, error) {

	// loadTokenFromCache the configuration from client_secret.json
	config, err := oa.config()
	if err != nil {
		return nil, err
	}
	// loadTokenFromCache the token from the cache or force authentication
	token, err := oa.loadTokenFromCache()
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, errors.New("not authenticated yet! please use birthday-pal oauth")
	}

	// Create the API client with a background context.
	ctx := context.Background()
	client := config.Client(ctx, token)

	return client, nil
}

// saveTokenInCache the token to the specified path (and saveTokenInCache the path to the struct).
// If the path is empty, then it will saveTokenInCache the path to the current cachePath.
func (oa OAuth2) saveTokenInCache(token *oauth2.Token) error {

	path := oa.system.CachePath()

	// Open the file for writing
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer f.Close()

	// Encode the token and write to disk
	if err := json.NewEncoder(f).Encode(token); err != nil {
		return fmt.Errorf("could not encode oauth token: %v", err)
	}

	return nil
}

// loadTokenFromCache the token from the specified path (and saves the path to the struct).
// If path is an empty string then it will loadTokenFromCache the token from the default
// cache path in the home directory. This method returns an error if the token
// cannot be loaded from the file.
func (oa OAuth2) loadTokenFromCache() (*oauth2.Token, error) {

	path := oa.system.CachePath()

	// Open the file at the path
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open cache file at %s: %v", path, err)
	}
	defer f.Close()

	// Decode the JSON token cache
	token := new(oauth2.Token)
	if err := json.NewDecoder(f).Decode(token); err != nil {
		return nil, fmt.Errorf("could not decode token in cache file at %s: %v", path, err)
	}

	return token, nil
}

//Authenticate performs a an OAuth2 authentication and save token in cache
func (oa OAuth2) Authenticate() error {

	config, err := oa.config()
	if err != nil {
		return err
	}

	// Compute the URL for the authorization
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	// Notify the user of the web browser.
	fmt.Println("Please use a browser in order to let birthday-pal access your contacts.")

	// Open the web browser
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", authURL).Start()
	case "windows", "darwin":
		err = exec.Command("open", authURL).Start()
	default:
		err = fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	// If we couldn't open the web browser, prompt the user to do it manually.
	if err != nil {
		fmt.Printf("Copy and paste the following link: \n%s\n\n", authURL)
	}

	// prompt for the authorization code
	code, err := oa.system.Prompt()
	if err != nil {
		return fmt.Errorf("unable to read authorization code %v", err)
	}

	// Perform the exchange for the token
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web %v", err)
	}

	// Cache the token to disk
	oa.saveTokenInCache(token)

	return nil
}

// config loads the secret file from the SecretPath. It is used both to
// create the client for requests as well as to perform authentication.
func (oa OAuth2) config() (*oauth2.Config, error) {

	data, err := ioutil.ReadFile(oa.SecretPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(data, oa.Scope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	return config, nil
}

// cachePath computes the path to the credential token file, creating the
// directory if necessary and stores it in the authentication struct.
/*func (oa OAuth2) cache() (string, error) {
	// Get the user to look up the user's home directory
	path := oa.cachePath
	if path == "" {
		path = oa.defaultCachePath()
	}

	return path, nil
}*/
