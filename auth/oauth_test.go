package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func Test_should_return_oauth2_authenticated_client(t *testing.T) {

	jsonConfig :=
		`
	{
		"installed": {
			"client_id": "c0nf1d3ential",
			"redirect_uris": ["http://uri"]
		}
	}
`
	token := `{"access_token":"s3cr3t"}`

	tempDir := tempDir()
	defer os.RemoveAll(tempDir)
	tempFileConfig := tempFileWithName(jsonConfig, tempDir, "config")
	tempFileToken := tempFileWithName(token, tempDir, "token")

	sys := new(fakeSystem)
	sys.On("CachePath").Return(tempFileToken)

	oauth2 := OAuth2{SecretPath: tempFileConfig, system: sys}

	clt, err := oauth2.Client()

	assert.NoError(t, err)
	assert.NotEmpty(t, clt)
	sys.AssertExpectations(t)
}

func Test_should_get_token_from_cache(t *testing.T) {

	jsonToken := `{"access_token":"s3cr3t"}`

	tempDir := tempDir()
	defer os.RemoveAll(tempDir)
	tempFile := tempFile(jsonToken, tempDir)

	sys := new(fakeSystem)
	sys.On("CachePath").Return(tempFile)

	auth := OAuth2{system: sys}

	token, err := auth.loadTokenFromCache()

	assert.NoError(t, err)
	assert.Equal(t, "s3cr3t", token.AccessToken)
	sys.AssertExpectations(t)
}

func Test_should_get_config_from_client_secret_file(t *testing.T) {

	jsonConfig :=
		`
	{
		"installed": {
			"client_id": "c0nf1d3ential",
			"redirect_uris": ["http://uri"]
		}
	}
`

	tempDir := tempDir()
	defer os.RemoveAll(tempDir)
	tempFile := tempFile(jsonConfig, tempDir)

	auth := OAuth2{SecretPath: tempFile}

	token, err := auth.config()

	assert.NoError(t, err)
	assert.Equal(t, "c0nf1d3ential", token.ClientID)
	assert.Equal(t, "http://uri", token.RedirectURL)
}
/*
func Test_should_authenticate_with_config(t *testing.T) {

	jsonConfig :=
		`
	{
		"installed": {
			"client_id": "c0nf1d3ential",
			"redirect_uris": ["http://uri"]
		}
	}
`
	sys := new(fakeSystem)
	sys.On("Prompt").Return("yolo", nil)

	tempDir := tempDir()
	defer os.RemoveAll(tempDir)
	tempFile := tempFile(jsonConfig, tempDir)

	auth := OAuth2{SecretPath: tempFile, system: sys}

	err := auth.Authenticate()

	assert.NoError(t, err)
	sys.AssertExpectations(t)
}
*/
func Test_config_should_throw_error_when_client_secret_file_is_not_valid_json(t *testing.T) {

	jsonConfig := `{"invalid:}`

	tempDir := tempDir()
	defer os.RemoveAll(tempDir)
	tempFile := tempFile(jsonConfig, tempDir)

	auth := OAuth2{SecretPath: tempFile}

	config, err := auth.config()

	assert.Error(t, err)
	assert.Empty(t, config)
}

func Test_config_should_throw_error_when_client_secret_file_does_not_exist(t *testing.T) {

	auth := OAuth2{SecretPath: "unknown"}

	config, err := auth.config()

	assert.Error(t, err)
	assert.Empty(t, config)
}

func tempFile(content string, dir string) string {
	return tempFileWithName(content, dir, "tmp-file")
}

func tempFileWithName(content string, dir string, filename string) string {
	byteContent := []byte(content)
	tmpfn := filepath.Join(dir, filename)
	if err := ioutil.WriteFile(tmpfn, byteContent, 0666); err != nil {
		log.Fatal(err)
	}
	return tmpfn
}

func tempDir() string {
	dir, err := ioutil.TempDir("", "tmp-dir")
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

type fakeSystem struct {
	mock.Mock
}

func (fs *fakeSystem) Prompt() (string, error) {
	called := fs.Called()
	return called.String(0), called.Error(1)
}

func (fs *fakeSystem) CachePath() string {
	called := fs.Called()
	return called.String(0)
}
