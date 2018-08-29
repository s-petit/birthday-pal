package email

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"github.com/flashmob/go-guerrilla"
	"fmt"
	"github.com/s-petit/birthday-pal/vcardparser"
	"github.com/flashmob/go-guerrilla/log"
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

func Test_hostPort(t *testing.T) {
	sender := SMTPSender{"localhost", "2525", "", ""}
	hostPort := sender.hostPort()
	assert.Equal(t, "localhost:2525", hostPort)
}

func Test_send(t *testing.T) {
	d := startSmtpServer()

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	sender := SMTPSender{"localhost", "2525", "", ""}
	contact := vcardparser.RemindContact{"ttf", birthday, 12}
	error := sender.Send(contact, []string{"recipient@test", "recipient2@test"})
	assert.NoError(t, error)
	d.Shutdown()
}

func Test_send_error(t *testing.T) {
	d := startSmtpServer()

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	sender := SMTPSender{"localhost", "2525", "", ""}
	contact := vcardparser.RemindContact{"ttf", birthday, 12}
	error := sender.Send(contact, []string{"recipient@test2"})
	assert.EqualError(t, error, "454 4.1.1 Error: Relay access denied: test2")
	d.Shutdown()
}

func startSmtpServer() (guerrilla.Daemon) {
	cfg := &guerrilla.AppConfig{
		LogFile: log.OutputOff.String(),
		AllowedHosts: []string{"test"},
	}

	d := guerrilla.Daemon{Config: cfg}

	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
	}

	return d
}
