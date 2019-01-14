package cli

import (
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/request"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type fakeBirthdayPal struct {
	mock.Mock
}

func (fbp fakeBirthdayPal) Exec(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Reminder, recipients []string) error {
	args := fbp.Called(contactsProvider, smtp, reminder, recipients)
	return args.Error(0)
}

func Test_carddav(t *testing.T) {

	os.Args = []string{"",
		"--smtp-host=localhost",
		"--smtp-port=2525",
		"--smtp-user=user@test",
		"--smtp-pass=smtp-pass",
		"--days-before=3",
		"--lang=FR",
		"--remind-everyday",
		"carddav",
		"--user=login",
		"--pass=password",
		"--url=http://carddav",
		"recipient@test",
	}

	now := time.Now()

	bpal := new(fakeBirthdayPal)
	system := new(testdata.FakeSystem)
	system.On("Now").Return(now)

	expectedContactProvider := request.CardDavContactsProvider{AuthClient: auth.BasicAuth{Username: "login", Password: "password"}, URL: "http://carddav"}
	expectedSMTP := email.SMTPClient{Host: "localhost", Port: 2525, Username: "user@test", Password: "smtp-pass", Language: "FR"}
	expectedReminder := remind.Reminder{CurrentDate: now, InNbDays: 3, Inclusive: true}
	expectedRecipients := []string{"recipient@test"}

	bpal.On("Exec", expectedContactProvider, expectedSMTP, expectedReminder, expectedRecipients).Return(nil)

	Mowcli(bpal, system)
	bpal.AssertExpectations(t)
	system.AssertExpectations(t)

}

func Test_google(t *testing.T) {
	os.Args = []string{"",
		"--smtp-host=localhost",
		"--smtp-port=2525",
		"--smtp-user=user@test",
		"--smtp-pass=smtp-pass",
		"--days-before=3",
		"--remind-everyday",
		"google",
		"--url=http://google",
		"myProfile",
		"recipient@test",
	}

	now := time.Now()

	bpal := new(fakeBirthdayPal)
	system := new(testdata.FakeSystem)
	system.On("Now").Return(now)

	expectedContactProvider := request.GoogleContactsProvider{AuthClient: auth.OAuth2Authenticator{Scope: "https://www.googleapis.com/auth/contacts.readonly", Profile: auth.OAuthProfile{System: system, Profile: "myProfile"}}, URL: "http://google"}
	expectedSMTP := email.SMTPClient{Host: "localhost", Port: 2525, Username: "user@test", Password: "smtp-pass", Language: "EN"}
	expectedReminder := remind.Reminder{CurrentDate: now, InNbDays: 3, Inclusive: true}
	expectedRecipients := []string{"recipient@test"}

	bpal.On("Exec", expectedContactProvider, expectedSMTP, expectedReminder, expectedRecipients).Return(nil)

	Mowcli(bpal, system)
	bpal.AssertExpectations(t)
	system.AssertExpectations(t)
}

func Test_oauth_perform(t *testing.T) {

	jsonConfig := testdata.JsonOauthConfig("c0nf1d3ential")
	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	tempFile := testdata.TempFile(jsonConfig, tempDir)

	expectedOauthConfig := testdata.Oauth2Config("c0nf1d3ential")
	expectedOauthConfig.Scopes = []string{"https://www.googleapis.com/auth/contacts.readonly"}

	profile := "myProfile"
	os.Args = []string{"",
		"oauth",
		"perform",
		profile,
		tempFile,
	}

	bpal := new(fakeBirthdayPal)
	system := new(testdata.FakeSystem)

	system.On("Prompt").Return("yolo", nil)
	system.On("OpenBrowser", mock.Anything).Return(nil)
	system.On("ExchangeToken", expectedOauthConfig, "yolo").Return(&oauth2.Token{}, nil)
	system.On("HomeDir").Return(tempDir)

	Mowcli(bpal, system)

	bpal.AssertNotCalled(t, "Exec")
	bpal.AssertExpectations(t)
	system.AssertExpectations(t)
}

func Test_oauth_list(t *testing.T) {

	tempDir := testdata.TempDir()
	defer os.RemoveAll(tempDir)
	testdata.TempFileWithName("anyContent", filepath.Join(tempDir, auth.CacheDirectory, "john"), "token.json")

	os.Args = []string{"",
		"oauth",
		"list",
	}

	bpal := new(fakeBirthdayPal)
	system := new(testdata.FakeSystem)
	system.On("HomeDir").Return(tempDir)

	Mowcli(bpal, system)

	bpal.AssertNotCalled(t, "Exec")
	bpal.AssertExpectations(t)
	system.AssertExpectations(t)
}
