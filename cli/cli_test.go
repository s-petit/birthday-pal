package cli

import (
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/request"
	system2 "github.com/s-petit/birthday-pal/system"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
	"time"
)

type fakeBirthdayPal struct {
	mock.Mock
}

func (fbp fakeBirthdayPal) RemindBirthdays(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Reminder, recipients []string) error {
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
		"--remind-everyday",
		"carddav",
		"http://carddav",
		"recipient@test",
	}

	time := time.Now()

	bpal := new(fakeBirthdayPal)
	system := new(system2.FakeSystem)
	system.On("Now").Return(time)

	expectedContactProvider := request.CardDavContactsProvider{AuthClient: auth.BasicAuth{Username: "", Password: ""}, URL: "http://carddav"}
	expectedSMTP := email.SMTPClient{Host: "localhost", Port: 2525, Username: "user@test", Password: "smtp-pass"}
	expectedReminder := remind.Reminder{CurrentDate: time, NbDaysBeforeBDay: 3, EveryDayUntilBDay: true}
	expectedRecipients := []string{"recipient@test"}

	bpal.On("RemindBirthdays", expectedContactProvider, expectedSMTP, expectedReminder, expectedRecipients).Return(nil)

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
		"/path/secret.json",
		"http://google",
		"recipient@test",
	}

	time := time.Now()

	bpal := new(fakeBirthdayPal)
	system := new(system2.FakeSystem)
	system.On("Now").Return(time)

	expectedContactProvider := request.GoogleContactsProvider{AuthClient: auth.OAuth2{Scope: "https://www.googleapis.com/auth/contacts.readonly", SecretPath: "/path/secret.json", System: system}, URL: "http://google"}
	expectedSMTP := email.SMTPClient{Host: "localhost", Port: 2525, Username: "user@test", Password: "smtp-pass"}
	expectedReminder := remind.Reminder{CurrentDate: time, NbDaysBeforeBDay: 3, EveryDayUntilBDay: true}
	expectedRecipients := []string{"recipient@test"}

	bpal.On("RemindBirthdays", expectedContactProvider, expectedSMTP, expectedReminder, expectedRecipients).Return(nil)

	Mowcli(bpal, system)
	bpal.AssertExpectations(t)
	system.AssertExpectations(t)
}

func Test_oauth(t *testing.T) {

}
