package auth

import (
	"errors"
	"github.com/s-petit/birthday-pal/system"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"os"
	"testing"
)

func Test_should_return_oauth2_authenticated_client(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	token := `{"access_token":"s3cr3t"}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFileConfig := testdata.TempFileWithName(jsonConfig, tempDir, "config")
	tempFileToken := testdata.TempFileWithName(token, tempDir, "token")

	sys := new(system.FakeSystem)
	sys.On("CachePath").Return(tempFileToken)

	oauth2 := OAuth2{SecretPath: tempFileConfig, System: sys}

	clt, err := oauth2.Client()

	assert.NoError(t, err)
	assert.NotEmpty(t, clt)
	sys.AssertExpectations(t)
}

func Test_should_not_return_oauth2_client_when_oauth_config_not_found(t *testing.T) {

	oauth2 := OAuth2{SecretPath: ""}

	clt, err := oauth2.Client()

	assert.Error(t, err)
	assert.Empty(t, clt)
}

func Test_should_not_return_oauth2_client_when_authentication_token_not_found(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFileConfig := testdata.TempFileWithName(jsonConfig, tempDir, "config")

	sys := new(system.FakeSystem)
	sys.On("CachePath").Return("")

	oauth2 := OAuth2{SecretPath: tempFileConfig, System: sys}

	clt, err := oauth2.Client()

	assert.Error(t, err)
	assert.Empty(t, clt)
}

func Test_should_get_token_from_cache(t *testing.T) {

	jsonToken := `{"access_token":"s3cr3t"}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonToken, tempDir)

	sys := new(system.FakeSystem)
	sys.On("CachePath").Return(tempFile)

	auth := OAuth2{System: sys}

	token, err := auth.loadTokenFromCache()

	assert.NoError(t, err)
	assert.Equal(t, "s3cr3t", token.AccessToken)
	sys.AssertExpectations(t)
}

func Test_should_not_get_token_from_cache_when_json_not_deserilizable(t *testing.T) {

	jsonToken := `{{{}}}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonToken, tempDir)

	sys := new(system.FakeSystem)
	sys.On("CachePath").Return(tempFile)

	auth := OAuth2{System: sys}

	token, err := auth.loadTokenFromCache()

	assert.Error(t, err)
	assert.Empty(t, token)
	sys.AssertExpectations(t)
}

func Test_should_get_config_from_client_secret_file(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonConfig, tempDir)

	auth := OAuth2{SecretPath: tempFile}

	token, err := auth.config()

	assert.NoError(t, err)
	assert.Equal(t, "c0nf1d3ential", token.ClientID)
	assert.Equal(t, "http://uri", token.RedirectURL)
}

func Test_should_authenticate_with_config(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonConfig, tempDir)

	sys := new(system.FakeSystem)
	sys.On("Prompt").Return("yolo", nil)
	sys.On("OpenBrowser", mock.Anything).Return(nil)
	sys.On("ExchangeToken", oauth2Config("c0nf1d3ential"), "yolo").Return(&oauth2.Token{}, nil)
	sys.On("CachePath").Return(tempDir + "/cache-file")

	auth := OAuth2{SecretPath: tempFile, System: sys}

	err := auth.Authenticate()

	assert.NoError(t, err)
	sys.AssertExpectations(t)
}

func Test_should_not_authenticate_when_token_not_exchanged(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonConfig, tempDir)

	sys := new(system.FakeSystem)
	sys.On("Prompt").Return("yolo", nil)
	sys.On("OpenBrowser", mock.Anything).Return(nil)
	sys.On("ExchangeToken", oauth2Config("c0nf1d3ential"), "yolo").Return(&oauth2.Token{}, errors.New("oops"))

	auth := OAuth2{SecretPath: tempFile, System: sys}

	err := auth.Authenticate()

	assert.Error(t, err)
	sys.AssertExpectations(t)
}

func Test_should_not_authenticate_when_value_prompted_is_malformed(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonConfig, tempDir)

	sys := new(system.FakeSystem)
	sys.On("Prompt").Return("", errors.New("oops"))
	sys.On("OpenBrowser", mock.Anything).Return(nil)

	auth := OAuth2{SecretPath: tempFile, System: sys}

	err := auth.Authenticate()

	assert.Error(t, err)
	sys.AssertExpectations(t)
}

func Test_should_authenticate_even_when_browser_not_openable(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonConfig, tempDir)

	sys := new(system.FakeSystem)
	sys.On("Prompt").Return("yolo", nil)
	sys.On("OpenBrowser", mock.Anything).Return(errors.New("erf"))
	sys.On("ExchangeToken", &oauth2.Config{ClientID: "c0nf1d3ential", ClientSecret: "", Endpoint: oauth2.Endpoint{AuthURL: "", TokenURL: ""}, RedirectURL: "http://uri", Scopes: []string{""}}, "yolo").Return(&oauth2.Token{}, nil)
	sys.On("CachePath").Return(tempDir + "/cache-file")

	auth := OAuth2{SecretPath: tempFile, System: sys}

	err := auth.Authenticate()

	assert.NoError(t, err)
	sys.AssertExpectations(t)
}

func Test_should_not_authenticate_when_config_not_valid(t *testing.T) {

	jsonConfig := "{{{}"

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonConfig, tempDir)

	auth := OAuth2{SecretPath: tempFile}

	err := auth.Authenticate()

	assert.Error(t, err)
}

func Test_config_should_throw_error_when_client_secret_file_is_not_valid_json(t *testing.T) {

	jsonConfig := `{"invalid:}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonConfig, tempDir)

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

func Test_should_save_token_in_cache_then_load_it(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	sys := new(system.FakeSystem)
	sys.On("CachePath").Return(tempDir + "/token.json")

	auth := OAuth2{System: sys}

	tok := &oauth2.Token{AccessToken: "s3cr3t"}

	err := auth.saveTokenInCache(tok)

	assert.NoError(t, err)
	sys.AssertExpectations(t)

	cachedToken, err := auth.loadTokenFromCache()

	assert.NoError(t, err)
	assert.Equal(t, "s3cr3t", cachedToken.AccessToken)
	sys.AssertExpectations(t)
}

func Test_should_not_save_token_in_not_authorized_path(t *testing.T) {

	sys := new(system.FakeSystem)
	sys.On("CachePath").Return("/root/token.json")

	auth := OAuth2{System: sys}

	tok := &oauth2.Token{AccessToken: "s3cr3t"}

	err := auth.saveTokenInCache(tok)

	assert.Error(t, err)
	sys.AssertExpectations(t)
}

func oauth2Config(clientID string) *oauth2.Config {
	return &oauth2.Config{ClientID: clientID, ClientSecret: "", Endpoint: oauth2.Endpoint{AuthURL: "", TokenURL: ""}, RedirectURL: "http://uri", Scopes: []string{""}}
}
