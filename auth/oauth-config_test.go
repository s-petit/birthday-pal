package auth

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func Test_should_get_token_from_cache(t *testing.T) {

	tok := `
{"access_token":"lol"}

`

	content := []byte(tok)
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	auth := authentication{cachePath: tmpfn}

	token, err := auth.getToken()

	assert.NoError(t, err)
	assert.Equal(t, "lol", token.AccessToken)
}
