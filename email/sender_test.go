package email

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_formatFrenchDate(t *testing.T) {
	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	formattedDate := formatFrenchDate(birthday)
	assert.Equal(t, "22/08", formattedDate)
}

func Test_formatEnglishDate(t *testing.T) {
	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	formattedDate := formatEnglishDate(birthday)
	assert.Equal(t, "08/22", formattedDate)
}

/*
func Test_send_mail(t *testing.T) {
	sendMails()
}*/
