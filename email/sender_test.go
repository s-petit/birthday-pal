package email

import (
	"fmt"
	"github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/log"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_hostPort(t *testing.T) {
	sender := SMTPSender{"localhost", "2525", "", ""}
	hostPort := sender.hostPort()
	assert.Equal(t, "localhost:2525", hostPort)
}

func Test_send(t *testing.T) {
	d := startSMTPServer()

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	sender := SMTPSender{"localhost", "2525", "", ""}
	c := contact.Contact{Name: "ttf", BirthDate: birthday}
	e := sender.Send(c, []string{"recipient@test", "recipient2@test"})
	assert.NoError(t, e)
	d.Shutdown()
}

func Test_send_error(t *testing.T) {
	d := startSMTPServer()

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	sender := SMTPSender{"localhost", "2525", "", ""}
	c := contact.Contact{Name: "ttf", BirthDate: birthday}
	e := sender.Send(c, []string{"recipient@test2"})
	assert.EqualError(t, e, "454 4.1.1 Error: Relay access denied: test2")
	d.Shutdown()
}

func startSMTPServer() guerrilla.Daemon {
	cfg := &guerrilla.AppConfig{
		LogFile:      log.OutputOff.String(),
		AllowedHosts: []string{"test"},
	}

	d := guerrilla.Daemon{Config: cfg}

	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
	}

	return d
}
