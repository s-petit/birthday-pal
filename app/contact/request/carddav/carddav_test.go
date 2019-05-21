package carddav

import (
	"fmt"
	"github.com/s-petit/birthday-pal/app/contact"
	"github.com/s-petit/birthday-pal/app/contact/auth/basic"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	vcardContact = `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
BDAY:19831228
END:VCARD
`
)

func basicAuthCarddavServer() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok && username == "user" && password == "pass" {
			io.WriteString(w, vcardContact)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
			w.WriteHeader(401)
		}
	}

	r := http.NewServeMux()
	r.HandleFunc("/contact", h)
	return r
}

func Test_GetContacts_should_return_carddav_contacts(t *testing.T) {

	srv := httptest.NewServer(basicAuthCarddavServer())
	defer srv.Close()

	carddav := ContactsProvider{
		AuthClient: basic.Auth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/contact", srv.URL),
	}

	contacts, err := carddav.GetContacts()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(contacts))
	assert.Equal(t, contact.Contact{Name: "Alexis Foo", BirthDate: testdata.BirthDate(1983, time.December, 28)}, contacts[0])

}

func Test_call_should_return_carddav_payload(t *testing.T) {

	srv := httptest.NewServer(basicAuthCarddavServer())
	defer srv.Close()

	carddav := ContactsProvider{
		AuthClient: basic.Auth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/contact", srv.URL),
	}

	payload, err := carddav.call()

	assert.NoError(t, err)
	assert.Equal(t, vcardContact, payload)
}

func Test_call_carddav_should_return_error_when_url_goes_to_404(t *testing.T) {

	srv := httptest.NewServer(basicAuthCarddavServer())
	defer srv.Close()

	r := ContactsProvider{
		AuthClient: basic.Auth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/unknown", srv.URL),
	}

	payload, err := r.call()

	assert.Error(t, err)
	assert.Equal(t, "", payload)
}

func Test_call_carddav_should_return_error_when_url_is_malformed(t *testing.T) {

	srv := httptest.NewServer(basicAuthCarddavServer())
	defer srv.Close()

	r := ContactsProvider{
		AuthClient: basic.Auth{Username: "user", Password: "pass"},
		URL:        "http://://",
	}

	payload, err := r.call()

	assert.Error(t, err)
	assert.Equal(t, "", payload)
}
