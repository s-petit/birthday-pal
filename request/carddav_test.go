package request

import (
	"testing"
	"net/http/httptest"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"io"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
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

func handler() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, vcardContact)
	}

	r := http.NewServeMux()
	r.HandleFunc("/contact", h)
	return r
}

func Test_GetContacts_should_return_carddav_contacts(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	carddav := CardDavContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/contact", srv.URL),
	}

	contacts, err := carddav.GetContacts()

	assert.Equal(t, 1, len(contacts))
	assert.Equal(t, contact.Contact{"Alexis Foo", testdata.BirthDate(1983, time.December, 28)}, contacts[0])

	assert.NoError(t, err)
}

func Test_call_should_return_carddav_payload(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	carddav := CardDavContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/contact", srv.URL),
	}

	payload, err := carddav.call()

	assert.Equal(t, vcardContact, payload)
	assert.NoError(t, err)
}

func Test_call_carddav_should_return_error_when_url_goes_to_404(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := CardDavContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/unknown", srv.URL),
	}

	payload, err := r.call()

	assert.Equal(t, "", payload)
	assert.Error(t, err)
}

func Test_call_carddav_should_return_error_when_url_is_malformed(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := CardDavContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        "http://://",
	}

	payload, err := r.call()

	assert.Equal(t, "", payload)
	assert.Error(t, err)
}

