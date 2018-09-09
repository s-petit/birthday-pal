package google

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// Authentication is a wrapper for the OAuth token authentication process,
// reading and writing the token from a credentials file on disk. It is meant
// to perform 100% of the workflow of authentication.
type Authentication struct {
	cachePath  string        // the path to the token cache file on disk
	configPath string        // the path to the client_secret.json file on disk
	token      *oauth2.Token // the oauth token cached in the cache file
	Scope      string        // the oauth desired scope which delimits access to user data
}

// Token returns the authentication token by first looking on disk for the
// cache, and if it doesn't exist by executing the authentication.
func (auth *Authentication) Token() (*oauth2.Token, error) {
	// Load the token from the cache if it doesn't exist.
	if auth.token == nil {
		if err := auth.Load(""); err != nil {
			// If we cannot load the token from disk, authenticate
			if err != nil {
				if err = auth.Authenticate(); err != nil {
					return nil, err
				}
			}
		}
	}

	// Return the token
	return auth.token, nil
}

// Authenticate runs an interactive authentication on the command line,
// prompting the user to open a brower page and enter an authorization code.
// It will then fetch an token via OAuth and cache it as credentials. Note
// that this method will overwrite any previously cached token.
func (auth *Authentication) Authenticate() error {
	// Load and create the OAuth2 Configuration
	config, err := auth.Config()
	if err != nil {
		return err
	}

	// Compute the URL for the authorization
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	// Notify the user of the web browser.
	fmt.Println("In order to authenticate, use a browser to authorize birthday-pal")

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

	// Prompt for the authorization code
	code, err := Prompt("enter authorization code:")
	if err != nil {
		return fmt.Errorf("unable to read authorization code %v", err)
	}

	// Perform the exchange for the token
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web %v", err)
	}

	// Cache the token to disk
	auth.token = token
	auth.Save("")

	return nil
}

// Config loads the client_secret.json from the ConfigPath. It is used both to
// create the client for requests as well as to perform authentication.
func (auth *Authentication) Config() (*oauth2.Config, error) {
	path, err := auth.ConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(data, auth.Scope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	return config, nil
}

// CachePath computes the path to the credential token file, creating the
// directory if necessary and stores it in the authentication struct.
func (auth *Authentication) CachePath() (string, error) {
	if auth.cachePath == "" {
		// Get the user to look up the user's home directory
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		// Get the hidden credentials directory, making sure it's created
		cacheDir := filepath.Join(usr.HomeDir, ".birthday-pal")
		os.MkdirAll(cacheDir, 0700)

		// Determine the path to the token cache file
		cacheFile := url.QueryEscape("credentials.json")
		auth.cachePath = filepath.Join(cacheDir, cacheFile)
	}

	fmt.Println(auth.cachePath)

	return auth.cachePath, nil
}

// ConfigPath computes the path to the configuration file, client_secret.json.
func (auth *Authentication) ConfigPath() (string, error) {
	if auth.configPath == "" {
		// Get the user to look up the user's home directory
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		// Create the path to the default configuration
		auth.configPath = filepath.Join(
			usr.HomeDir, "dev/go/src/github.com/s-petit/birthday-pal/contact/google/.birthday-pal", "client_secret.json",
		)
	}

	return auth.configPath, nil
}

// Load the token from the specified path (and saves the path to the struct).
// If path is an empty string then it will load the token from the default
// cache path in the home directory. This method returns an error if the token
// cannot be loaded from the file.
func (auth *Authentication) Load(path string) error {
	var err error

	// Get the default cache path or save the specified path
	if path == "" {
		path, err = auth.CachePath()
		if err != nil {
			return err
		}
	} else {
		auth.cachePath = path
	}

	// Open the file at the path
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open cache file at %s: %v", path, err)
	}
	defer f.Close()

	// Decode the JSON token cache
	auth.token = new(oauth2.Token)
	if err := json.NewDecoder(f).Decode(auth.token); err != nil {
		return fmt.Errorf("could not decode token in cache file at %s: %v", path, err)
	}

	return nil
}

// Save the token to the specified path (and save the path to the struct).
// If the path is empty, then it will save the path to the current CachePath.
func (auth *Authentication) Save(path string) error {
	var err error

	// Get the default cache path or save the specified path
	if path == "" {
		path, err = auth.CachePath()
		if err != nil {
			return err
		}
	} else {
		auth.cachePath = path
	}

	// Open the file for writing
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer f.Close()

	// Encode the token and write to disk
	if err := json.NewEncoder(f).Encode(auth.token); err != nil {
		return fmt.Errorf("could not encode oauth token: %v", err)
	}

	return nil
}

// Delete the token file at the given path in order to force a
// reauthentication. This method also saves the given path, or if the path is
// empty, then it will compute the default CachePath. This method will not
// return an error on failure (e.g. if the file does not exist).
func (auth *Authentication) Delete(path string) {
	// Get the default cache path or save the specified path
	if path == "" {
		path, _ = auth.CachePath()
	} else {
		auth.cachePath = path
	}

	//  Delete the file at the cache path if it exists
	os.Remove(path)
}

func Prompt(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s ", prompt)

	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	response = strings.TrimSpace(response)
	if response == "" {
		return Prompt(prompt)
	}

	return response, nil
}
