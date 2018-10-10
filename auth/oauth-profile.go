package auth

import (
	"encoding/json"
	"fmt"
	"github.com/s-petit/birthday-pal/system"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
	"path/filepath"
)

//CacheDirectory is the location of bpal's cached files
const CacheDirectory = ".birthday-pal"

//OAuthProfile holds logic of cache storage of oauth authentication profiles
type OAuthProfile struct {
	System  system.System
	Profile string
}

//ListProfiles lists the registered profiles by looking inside cache directory
func (oap OAuthProfile) ListProfiles() ([]string, error) {
	cachePath := oap.cachePath()
	files, err := ioutil.ReadDir(cachePath)
	if err != nil {
		return nil, err
	}

	var profiles []string

	for _, f := range files {
		if f.IsDir() {
			profiles = append(profiles, f.Name())
		}
	}
	return profiles, nil
}

//cachePath is the location where all profiles will be stored in order to remember authentication.
func (oap OAuthProfile) cachePath() string {
	cacheDir := filepath.Join(oap.System.HomeDir(), CacheDirectory)
	return cacheDir
}

//profileCachePath is the cache location for a given profile.
func (oap OAuthProfile) profileCachePath() string {
	// Get the hidden credentials directory, making sure it's created
	cacheDir := filepath.Join(oap.cachePath(), oap.Profile)
	os.MkdirAll(cacheDir, 0700)
	return cacheDir
}

// save the token to the specified path.
func (oap OAuthProfile) saveTokenInCache(token *oauth2.Token) error {

	tokenPath := filepath.Join(oap.profileCachePath(), tokenFile)

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
func (oap OAuthProfile) saveConfigInCache(secretPath string) error {

	data, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %v", err)
	}

	conf := filepath.Join(oap.profileCachePath(), configFile)
	ioutil.WriteFile(conf, data, 0644)

	return nil
}

// load the oauth2 token from the specified path
func (oap OAuthProfile) loadTokenFromCache() (*oauth2.Token, error) {

	cachePath := oap.profileCachePath()
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
