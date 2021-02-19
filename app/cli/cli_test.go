package cli

import (
	"github.com/s-petit/birthday-pal/app/contact/auth"
	"github.com/s-petit/birthday-pal/app/contact/request"
	"github.com/s-petit/birthday-pal/app/email"
	"github.com/s-petit/birthday-pal/app/remind"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
	"time"
)

type fakeBirthdayPal struct {
	mock.Mock
}

func (fbp fakeBirthdayPal) Exec(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Params, recipients []string) error {
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
		"--less-than-or-equal",
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
	expectedReminder := remind.Params{Today: now, InNbDays: 3, Inclusive: true}
	expectedRecipients := []string{"recipient@test"}

	bpal.On("Exec", expectedContactProvider, expectedSMTP, expectedReminder, expectedRecipients).Return(nil)

	Mowcli(bpal, system)
	bpal.AssertExpectations(t)
	system.AssertExpectations(t)

}
