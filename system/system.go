package system

import (
	"bufio"
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// System holds system-dependant methods which are hard to test/mock
type System interface {
	CachePath() string
	Prompt() (string, error)
	OpenBrowser(url string) error
	ExchangeToken(config *oauth2.Config, code string) (*oauth2.Token, error)
}

// RealSystem is how the hosting system works in real life
type RealSystem struct {
}

//CachePath is the location where the token will be stored in order to remember authentication.
func (rs RealSystem) CachePath() string {
	// Get the hidden credentials directory, making sure it's created
	cacheDir := filepath.Join(homeDir(), ".birthday-pal")
	os.MkdirAll(cacheDir, 0700)
	// Determine the path to the token cache file
	cacheFile := url.QueryEscape("credentials.json")
	return filepath.Join(cacheDir, cacheFile)
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

//Prompt asks user the auth code returned by oauth server.
func (rs RealSystem) Prompt() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	promptMsg := "enter authorization code:"
	fmt.Printf("%s ", promptMsg)

	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	response = strings.TrimSpace(response)
	if response == "" {
		return rs.Prompt()
	}

	return response, nil
}

//OpenBrowser opens a browser in order to login and authorize bpal to use oauth2 secured apis.
func (rs RealSystem) OpenBrowser(authURL string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", authURL).Start()
	case "windows", "darwin":
		return exec.Command("open", authURL).Start()
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

//ExchangeToken calls google auth server returns a oauth2 token if authentication succeeded
func (rs RealSystem) ExchangeToken(config *oauth2.Config, code string) (*oauth2.Token, error) {
	return config.Exchange(context.Background(), code)
}
