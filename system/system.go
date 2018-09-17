package system

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type System interface {
	CachePath() string
	Prompt() (string, error)
}

type RealSystem struct {
}

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

/*func Test_should_get_default_cache_path(t *testing.T) {
	oauth2 := OAuth2{}
	path := oauth2.defaultCachePath()
	assert.Contains(t, path, "/.birthday-pal/credentials.json")
}
*/
/*func Test_should_get_overriden_cache_path(t *testing.T) {

	sys := new(fakeSystem)
	sys.On("HomeDir").Return(tempDir)


	auth := OAuth2{system: sys}
	path, err := auth.cache()

	assert.NoError(t, err)
	assert.Contains(t, path, "/tmp/creds.json")
	sys.AssertExpectations(t)
}*/
