package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_should_return_basicauth_authenticated_client(t *testing.T) {

	basic := BasicAuth{Username: "user", Password: "pass"}

	client, err := basic.Client()
	assert.NoError(t, err)

	srv := httptest.NewServer(basicAuthServer())
	defer srv.Close()

	resp, e := client.Get(fmt.Sprintf("%s/myurl", srv.URL))
	assert.NoError(t, e)
	assert.Equal(t, 200, resp.StatusCode)

}

func basicAuthServer() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok && username == "user" && password == "pass" {
			w.WriteHeader(200)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
			w.WriteHeader(401)
		}
	}

	r := http.NewServeMux()
	r.HandleFunc("/myurl", h)
	return r
}
