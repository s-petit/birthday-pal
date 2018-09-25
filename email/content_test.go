package email

import (
	"github.com/s-petit/birthday-pal/remind"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_should_get_mail_in_french(t *testing.T) {

	expectedMail := `
To: Birthday Pals\r\n
Subject: Anniversaire de John -34 an(s)-\r\n
Ce sera l'anniversaire de John le 22/08. Pensez à le lui souhaiter!
`
	birthday := time.Date(1980, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, _ := toMail(remind.ContactBirthday{"John", birthday, 34}, "fr")
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_french_without_age(t *testing.T) {

	expectedMail := `
To: Birthday Pals\r\n
Subject: Anniversaire de John\r\n
Ce sera l'anniversaire de John le 22/08. Pensez à le lui souhaiter!
`
	birthday := time.Date(0, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, _ := toMail(remind.ContactBirthday{"John", birthday, 34}, "fr")
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_english(t *testing.T) {

	expectedMail := `
To: Birthday Pals\r\n
Subject: John's birthday -34 yo-\r\n
The 08/22 will be John's birthday. Do not forget to make your wish!
`
	birthday := time.Date(1980, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, _ := toMail(remind.ContactBirthday{"John", birthday, 34}, "EN")
	assert.Equal(t, expectedMail, string(bytes))
}
