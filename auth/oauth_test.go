package auth

import (
	"errors"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"os"
	"path/filepath"
	"testing"
)

func Test_should_return_oauth2_authenticated_client(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	token := `{"access_token":"s3cr3t"}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	testdata.TempFileWithName(jsonConfig, filepath.Join(tempDir, CacheDirectory), "config.json")
	testdata.TempFileWithName(token, filepath.Join(tempDir, CacheDirectory), "token.json")

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	authenticator := OAuth2Authenticator{Profile: OAuthProfile{System: sys}}

	clt, err := authenticator.Client()

	assert.NoError(t, err)
	assert.NotEmpty(t, clt)
	sys.AssertExpectations(t)
}

func Test_should_not_return_oauth2_client_when_oauth_config_not_found(t *testing.T) {
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	authenticator := OAuth2Authenticator{Profile: OAuthProfile{System: sys}}

	clt, err := authenticator.Client()

	assert.Error(t, err)
	assert.Empty(t, clt)
}

func Test_should_not_return_oauth2_client_when_authentication_token_not_found(t *testing.T) {

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return("")

	authenticator := OAuth2Authenticator{Profile: OAuthProfile{System: sys}}

	clt, err := authenticator.Client()

	assert.Error(t, err)
	assert.Empty(t, clt)
}

func Test_should_get_config_from_client_secret_file(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	profileName := "authProfile"
	testdata.TempFileWithName(jsonConfig, filepath.Join(tempDir, CacheDirectory, profileName), "config.json")

	sys := new(testdata.FakeSystem)
	profile := OAuthProfile{Profile: profileName, System: sys}

	sys.On("HomeDir").Return(tempDir)

	auth := OAuth2Authenticator{Profile: profile}

	token, err := auth.config()

	assert.NoError(t, err)
	assert.Equal(t, "c0nf1d3ential", token.ClientID)
	assert.Equal(t, "http://uri", token.RedirectURL)
}

func Test_should_authenticate_with_config(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	profileName := "authProfile"
	secretPath := testdata.TempFileWithName(jsonConfig, filepath.Join(tempDir, CacheDirectory, profileName), "config.json")

	sys := new(testdata.FakeSystem)
	profile := OAuthProfile{Profile: profileName, System: sys}

	sys.On("Prompt").Return("yolo", nil)
	sys.On("OpenBrowser", mock.Anything).Return(nil)
	sys.On("ExchangeToken", testdata.Oauth2Config("c0nf1d3ential"), "yolo").Return(&oauth2.Token{}, nil)
	sys.On("HomeDir").Return(tempDir)

	auth := OAuth2Authenticator{Profile: profile}

	err := auth.Authenticate(secretPath)

	assert.NoError(t, err)
	sys.AssertExpectations(t)
}

func Test_should_not_authenticate_when_token_not_exchanged(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	profileName := "authProfile"
	tempFile := testdata.TempFile(jsonConfig, filepath.Join(tempDir, CacheDirectory, profileName))

	sys := new(testdata.FakeSystem)
	profile := OAuthProfile{Profile: profileName, System: sys}
	sys.On("Prompt").Return("yolo", nil)
	sys.On("OpenBrowser", mock.Anything).Return(nil)
	sys.On("ExchangeToken", testdata.Oauth2Config("c0nf1d3ential"), "yolo").Return(&oauth2.Token{}, errors.New("oops"))
	sys.On("HomeDir").Return(tempDir)

	auth := OAuth2Authenticator{Profile: profile}

	err := auth.Authenticate(tempFile)

	assert.Error(t, err)
	sys.AssertExpectations(t)
}

func Test_should_not_authenticate_when_value_prompted_is_malformed(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	profileName := "authProfile"
	tempFile := testdata.TempFile(jsonConfig, filepath.Join(tempDir, CacheDirectory, profileName))
	sys := new(testdata.FakeSystem)
	profile := OAuthProfile{Profile: profileName, System: sys}

	sys.On("Prompt").Return("", errors.New("oops"))
	sys.On("OpenBrowser", mock.Anything).Return(nil)
	sys.On("HomeDir").Return(tempDir)

	auth := OAuth2Authenticator{Profile: profile}

	err := auth.Authenticate(tempFile)

	assert.Error(t, err)
	sys.AssertExpectations(t)
}

func Test_should_authenticate_even_when_browser_not_openable(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	profileName := "authProfile"
	tempFile := testdata.TempFile(jsonConfig, filepath.Join(tempDir, CacheDirectory, profileName))
	sys := new(testdata.FakeSystem)

	profile := OAuthProfile{Profile: profileName, System: sys}

	sys.On("Prompt").Return("yolo", nil)
	sys.On("OpenBrowser", mock.Anything).Return(errors.New("erf"))
	sys.On("ExchangeToken", &oauth2.Config{ClientID: "c0nf1d3ential", ClientSecret: "", Endpoint: oauth2.Endpoint{AuthURL: "", TokenURL: ""}, RedirectURL: "http://uri", Scopes: []string{""}}, "yolo").Return(&oauth2.Token{}, nil)
	sys.On("HomeDir").Return(tempDir)

	auth := OAuth2Authenticator{Profile: profile}

	err := auth.Authenticate(tempFile)

	assert.NoError(t, err)
	sys.AssertExpectations(t)
}

func Test_should_not_authenticate_when_config_not_valid(t *testing.T) {

	jsonConfig := "{{{}"

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	profileName := "authProfile"
	tempFile := testdata.TempFile(jsonConfig, filepath.Join(tempDir, CacheDirectory, profileName))

	sys := new(testdata.FakeSystem)
	profile := OAuthProfile{Profile: profileName, System: sys}

	sys.On("HomeDir").Return(tempDir)

	auth := OAuth2Authenticator{Profile: profile}

	err := auth.Authenticate(tempFile)

	assert.Error(t, err)
}

func Test_config_should_throw_error_when_client_secret_file_is_not_valid_json(t *testing.T) {

	jsonConfig := `{"invalid:}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	profileName := "authProfile"
	testdata.TempFile(jsonConfig, filepath.Join(tempDir, CacheDirectory, profileName))

	sys := new(testdata.FakeSystem)
	profile := OAuthProfile{Profile: profileName, System: sys}

	sys.On("HomeDir").Return(tempDir)

	auth := OAuth2Authenticator{Profile: profile}

	config, err := auth.config()

	assert.Error(t, err)
	assert.Empty(t, config)
}

func Test_config_should_throw_error_when_client_secret_file_does_not_exist(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	sys := new(testdata.FakeSystem)
	profile := OAuthProfile{Profile: "authProfile", System: sys}

	sys.On("HomeDir").Return(tempDir)

	auth := OAuth2Authenticator{Profile: profile}

	config, err := auth.config()

	assert.Error(t, err)
	assert.Empty(t, config)
}
