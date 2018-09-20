package testdata

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"time"
)

//FakeSystem represents a mockable system
type FakeSystem struct {
	mock.Mock
}

//Now return a mocked now
func (fs *FakeSystem) Now() time.Time {
	called := fs.Called()
	return called.Get(0).(time.Time)
}

//Prompt returns a mocked prompt
func (fs *FakeSystem) Prompt() (string, error) {
	called := fs.Called()
	return called.String(0), called.Error(1)
}

//CachePath returns a mocked cache path
func (fs *FakeSystem) CachePath() string {
	called := fs.Called()
	return called.String(0)
}

//OpenBrowser mocks the browser opening
func (fs *FakeSystem) OpenBrowser(URL string) error {
	called := fs.Called(URL)
	return called.Error(0)
}

//ExchangeToken mocks the token exchange inside google auth server
func (fs *FakeSystem) ExchangeToken(config *oauth2.Config, code string) (*oauth2.Token, error) {
	called := fs.Called(config, code)
	return called.Get(0).(*oauth2.Token), called.Error(1)
}
