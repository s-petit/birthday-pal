package auth

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

/*func Test_should_return_basicauth_authenticated_client(t *testing.T) {

	basic := BasicAuth{Username: "user", Password: "pass"}

	_, err := basic.Client()

	assert.NoError(t, err)
}
*/

func Test_should_get_default_cache_path(t *testing.T) {
	path := defaultCachePath()
	assert.Contains(t, path, "/.birthday-pal/credentials.json")
}

func Test_should_get_overriden_cache_path(t *testing.T) {
	auth := OAuth2{cachePath: "/tmp/creds.json"}
	path, err := auth.cache()

	assert.NoError(t, err)
	assert.Contains(t, path, "/tmp/creds.json")
}

func Test_should_get_token_from_cache(t *testing.T) {

	tok := `{"access_token":"s3cr3t"}`

	content := []byte(tok)
	dir, err := ioutil.TempDir("", "tmp-dir")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmp-file")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	auth := OAuth2{cachePath: tmpfn}

	token, err := auth.loadTokenFromCache()

	assert.NoError(t, err)
	assert.Equal(t, "s3cr3t", token.AccessToken)
}


func Test_should_get_config_from_client_secret_file(t *testing.T) {

	tok :=
`
	{
		"installed": {
			"client_id": "c0nf1d3ential",
			"redirect_uris": ["http://uri"]
		}
	}
`

	content := []byte(tok)
	dir, err := ioutil.TempDir("", "tmp-dir")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmp-file")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	auth := OAuth2{SecretPath: tmpfn}

	token, err := auth.config()

	assert.NoError(t, err)
	assert.Equal(t, "c0nf1d3ential", token.ClientID)
	assert.Equal(t, "http://uri", token.RedirectURL)
}
