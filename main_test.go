package main

import (
	"fmt"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"
)

func handler() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		vcard := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
N:Foo;Alexis;;;
BDAY:%s
END:VCARD
`
		bday := vcardBday(time.Now())
		io.WriteString(w, fmt.Sprintf(vcard, bday))
	}

	r := http.NewServeMux()
	r.HandleFunc("/contact", h)
	return r
}

//Integration test
func Test_main_with_carddav(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	d := testdata.StartSMTPServer()
	defer d.Shutdown()

	os.Args = []string{"",
		"--smtp-host=localhost",
		"--smtp-port=2525",
		"carddav",
		fmt.Sprintf("--url=%s/contact", srv.URL),
		"recipient@test",
	}

	assert.NotPanics(t, main)
}

func vcardBday(date time.Time) string {
	year := strconv.Itoa(date.Year())
	month := strconv.Itoa(int(date.Month()))
	day := strconv.Itoa(date.Day())

	if int(date.Month()) < 10 {
		month = "0" + month
	}
	if date.Day() < 10 {
		day = "0" + day
	}
	bday := year + month + day
	return bday
}
