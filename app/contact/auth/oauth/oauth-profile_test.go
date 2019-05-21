package oauth

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"os"
	"path/filepath"
	"testing"
)

func Test_should_return_oauth_registered_profiles_inside_cache_path_as_directories(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	johnProfile := filepath.Join(tempDir, CacheDirectory, "john")
	os.MkdirAll(johnProfile, 0700)

	bobProfile := filepath.Join(tempDir, CacheDirectory, "bob")
	os.MkdirAll(bobProfile, 0700)

	testdata.TempFileWithName("any", filepath.Join(tempDir, CacheDirectory), "bob.json")

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	profiles, err := Profile{System: sys}.ListProfiles()

	assert.NoError(t, err)
	assert.Equal(t, []string{"bob", "john"}, profiles)
}

func Test_should_not_return_oauth_registered_profiles_when_cache_path_does_not_exist(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	profiles, err := Profile{System: sys}.ListProfiles()

	assert.Error(t, err)
	assert.Empty(t, profiles)
}

func Test_not_should_get_token_from_cache_when_token_not_found(t *testing.T) {
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	profileName := "authProfile"

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	auth := Profile{Profile: profileName, System: sys}

	token, err := auth.loadProfileTokenFromCache()

	assert.Error(t, err)
	assert.Empty(t, token)
	sys.AssertExpectations(t)
}

func Test_should_not_get_token_from_cache_when_json_not_deserilizable(t *testing.T) {

	jsonToken := `{{{}}}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	profileName := "authProfile"
	testdata.TempFileWithName(jsonToken, filepath.Join(tempDir, CacheDirectory, profileName), "token.json")

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	auth := Profile{Profile: profileName, System: sys}

	token, err := auth.loadProfileTokenFromCache()

	assert.Error(t, err)
	assert.Empty(t, token)
	sys.AssertExpectations(t)
}

func Test_should_save_token_in_cache_then_load_it(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	auth := Profile{Profile: "authProfile", System: sys}

	tok := &oauth2.Token{AccessToken: "s3cr3t"}

	err := auth.saveProfileTokenInCache(tok)

	assert.NoError(t, err)
	sys.AssertExpectations(t)

	cachedToken, err := auth.loadProfileTokenFromCache()

	assert.NoError(t, err)
	assert.Equal(t, "s3cr3t", cachedToken.AccessToken)
	sys.AssertExpectations(t)
}

func Test_should_not_save_token_in_not_authorized_path(t *testing.T) {

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return("/root")

	auth := Profile{Profile: "authProfile", System: sys}

	tok := &oauth2.Token{AccessToken: "s3cr3t"}

	err := auth.saveProfileTokenInCache(tok)

	assert.Error(t, err)
	sys.AssertExpectations(t)
}

func Test_config_should_throw_error_when_client_secret_file_is_not_valid_json(t *testing.T) {

	jsonConfig := `{"invalid:}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	profileName := "authProfile"
	testdata.TempFile(jsonConfig, filepath.Join(tempDir, CacheDirectory, profileName))

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	profile := Profile{Profile: profileName, System: sys}

	config, err := profile.loadProfileConfigFromCache("")

	assert.Error(t, err)
	assert.Empty(t, config)
}

func Test_config_should_throw_error_when_client_secret_file_does_not_exist(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)
	profile := Profile{Profile: "authProfile", System: sys}

	config, err := profile.loadProfileConfigFromCache("")

	assert.Error(t, err)
	assert.Empty(t, config)
}

func Test_should_save_config_from_client_secret_file_then_load_it(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	profileName := "authProfile"
	configPath := testdata.TempFileWithName(jsonConfig, filepath.Join(tempDir), "config.json")

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)
	profile := Profile{Profile: profileName, System: sys}

	profile.saveProfileConfigInCache(configPath)

	token, err := profile.loadProfileConfigFromCache("")

	assert.NoError(t, err)
	assert.Equal(t, "c0nf1d3ential", token.ClientID)
	assert.Equal(t, "http://uri", token.RedirectURL)
}

func Test_should_not_save_config_when_config_path_not_exists(t *testing.T) {

	profileName := "authProfile"

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return("any")
	profile := Profile{Profile: profileName, System: sys}

	profile.saveProfileConfigInCache("anypath")

	token, err := profile.loadProfileConfigFromCache("")

	assert.Error(t, err)
	assert.Empty(t, token)
}
