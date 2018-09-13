package request

import (
	"fmt"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	googleContact = `
{
	"connections": [
		{
      		"names": [
      		  {
      		    "displayName": "Alexis Foo"
      		  }
      		],
      		"birthdays": [
      		  {
      		    "date": {
      		      "year": 1983,
      		      "month": 12,
      		      "day": 28
      		    }
      		  }
      		]
    	}
	]
}
`
)

func googleHandler() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, googleContact)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"contact\": \"Alexis\"}")
	}

	h3 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{[}")
	}

	r := http.NewServeMux()
	r.HandleFunc("/contact", h)
	r.HandleFunc("/other-api", h2)
	r.HandleFunc("/not-json", h3)
	return r
}

func Test_GetContacts_should_return_google_contacts(t *testing.T) {

	srv := httptest.NewServer(googleHandler())
	defer srv.Close()

	google := GoogleContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/contact", srv.URL),
	}

	contacts, err := google.GetContacts()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(contacts))
	assert.Equal(t, contact.Contact{"Alexis Foo", testdata.BirthDate(1983, time.December, 28)}, contacts[0])

}

func Test_GetContacts_should_not_return_contact_when_payload_is_not_from_people_api(t *testing.T) {

	srv := httptest.NewServer(googleHandler())
	defer srv.Close()

	google := GoogleContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/other-api", srv.URL),
	}

	contacts, err := google.GetContacts()

	assert.NoError(t, err)
	assert.Equal(t, 0, len(contacts))
}

func Test_GetContacts_should_return_error_when_payload_is_not_valid_json(t *testing.T) {

	srv := httptest.NewServer(googleHandler())
	defer srv.Close()

	google := GoogleContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/not-json", srv.URL),
	}

	contacts, err := google.GetContacts()

	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
}

func Test_GetContacts_should_return_error_when_url_goes_to_404(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	google := GoogleContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/unknown", srv.URL),
	}

	contacts, err := google.GetContacts()

	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
}

func Test_GetContacts_should_return_error_when_url_is_malformed(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	google := GoogleContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        "http://://",
	}

	contacts, err := google.GetContacts()

	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
}
