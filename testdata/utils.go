package testdata

import (
	"fmt"
	"github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/log"
	"time"
)

func BirthDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func LocalDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func StartSMTPServer() guerrilla.Daemon {
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
