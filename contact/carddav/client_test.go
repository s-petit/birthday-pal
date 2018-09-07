package carddav

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		vcard := "myVCard"
		io.WriteString(w, vcard)
	}

	r := http.NewServeMux()
	r.HandleFunc("/contact", h)
	return r
}

func Test_Get_should_return_payload(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := BasicAuthRequest{fmt.Sprintf("%s/contact", srv.URL), "user", "pass"}

	s, err := r.Get()

	assert.Equal(t, "myVCard", s)
	assert.NoError(t, err)
}

func Test_Get_should_return_error_when_url_goes_to_404(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := BasicAuthRequest{fmt.Sprintf("%s/unknown", srv.URL), "user", "pass"}

	s, err := r.Get()

	assert.Equal(t, "", s)
	assert.Error(t, err)
}

func Test_Get_should_return_error_when_url_is_malformed(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := BasicAuthRequest{"http://://", "user", "pass"}

	s, err := r.Get()

	assert.Equal(t, "", s)
	assert.Error(t, err)
}
