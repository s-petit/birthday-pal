package auth

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"os"
	"path/filepath"
	"testing"
)

func Test_should_return_oauth_registered_profiles_inside_cache_path(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	johnProfile := filepath.Join(tempDir, CacheDirectory, "john")
	os.MkdirAll(johnProfile, 0700)

	bobProfile := filepath.Join(tempDir, CacheDirectory, "bob")
	os.MkdirAll(bobProfile, 0700)

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	profiles, err := OAuthProfile{System: sys}.ListProfiles()

	assert.NoError(t, err)
	assert.Equal(t, []string{"bob", "john"}, profiles)
}

func Test_should_get_token_from_cache(t *testing.T) {

	jsonToken := `{"access_token":"s3cr3t"}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	testdata.TempFileWithName(jsonToken, filepath.Join(tempDir, CacheDirectory, "authProfile"), "token.json")

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	auth := OAuthProfile{Profile: "authProfile", System: sys}

	token, err := auth.loadTokenFromCache()

	assert.NoError(t, err)
	assert.Equal(t, "s3cr3t", token.AccessToken)
	sys.AssertExpectations(t)
}

func Test_should_not_get_token_from_cache_when_json_not_deserilizable(t *testing.T) {

	jsonToken := `{{{}}}`

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	testdata.TempFileWithName(jsonToken, filepath.Join(tempDir, CacheDirectory), "token.json")

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	auth := OAuthProfile{Profile: "authProfile", System: sys}

	token, err := auth.loadTokenFromCache()

	assert.Error(t, err)
	assert.Empty(t, token)
	sys.AssertExpectations(t)
}

func Test_should_save_token_in_cache_then_load_it(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return(tempDir)

	auth := OAuthProfile{Profile: "authProfile", System: sys}

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

	sys := new(testdata.FakeSystem)
	sys.On("HomeDir").Return("/root")

	auth := OAuthProfile{Profile: "authProfile", System: sys}

	tok := &oauth2.Token{AccessToken: "s3cr3t"}

	err := auth.saveTokenInCache(tok)

	assert.Error(t, err)
	sys.AssertExpectations(t)
}
