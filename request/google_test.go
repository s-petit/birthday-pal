package request

import (
	"net/http"
	"io"
	"testing"
	"net/http/httptest"
	"github.com/s-petit/birthday-pal/auth"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"time"
)

var (
	googleContact = `
{
      "resourceName": "people/c2442229414628784193",
      "etag": "%EgYBAj0HNy4aDAECAwQFBgcICQoLDCIMS3RaMzZHN0s4MkE9",
      "names": [
        {
          "metadata": {
            "primary": true,
            "source": {
              "type": "CONTACT",
              "id": "21e48ab68ef57841"
            }
          },
          "displayName": "Naoki Francomme",
          "familyName": "Francomme",
          "givenName": "Naoki",
          "displayNameLastFirst": "Francomme, Naoki"
        }
      ],
      "birthdays": [
        {
          "metadata": {
            "primary": true,
            "source": {
              "type": "CONTACT",
              "id": "21e48ab68ef57841"
            }
          },
          "date": {
            "month": 11,
            "day": 15
          },
          "text": "15/11"
        }
      ]
    }
`
)

func googleHandler() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, googleContact)
	}

	r := http.NewServeMux()
	r.HandleFunc("/contact", h)
	return r
}

func Test_GetContacts_should_return_google_contacts(t *testing.T) {

	srv := httptest.NewServer(handler())
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
/*
func Test_call_should_return_google_payload(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	carddav := GoogleContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/contact", srv.URL),
	}

	payload, err := carddav.call()

	assert.Equal(t, googleContact, payload)
	assert.NoError(t, err)
}

func Test_call_google_should_return_error_when_url_goes_to_404(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := GoogleContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        fmt.Sprintf("%s/unknown", srv.URL),
	}

	payload, err := r.call()

	assert.Equal(t, "", payload)
	assert.Error(t, err)
}

func Test_call_google_should_return_error_when_url_is_malformed(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := GoogleContactsProvider{
		AuthClient: auth.BasicAuth{Username: "user", Password: "pass"},
		URL:        "http://://",
	}

	payload, err := r.call()

	assert.Equal(t, "", payload)
	assert.Error(t, err)
}


*/