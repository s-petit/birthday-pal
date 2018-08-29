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
	//r.HandleFunc("/nil", h2)
	return r
}

func Test_Ok(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := BasicAuthRequest{fmt.Sprintf("%s/contact", srv.URL), "user", "pass"}

	s, err := r.Get()

	assert.Equal(t, "myVCard", s)
	assert.NoError(t, err)
}

func Test_404(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	r := BasicAuthRequest{fmt.Sprintf("%s/unknown", srv.URL), "user", "pass"}

	s, err := r.Get()

	assert.Equal(t, "", s)
	assert.Error(t, err)
}

/*	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))*/
