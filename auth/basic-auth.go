package auth

import (
	"net/http"
	"net/url"
)

//BasicAuth provides
type BasicAuth struct {
	Username string
	Password string
}

//Client returns a HTTP client authenticated with basic Auth
func (ba BasicAuth) Client() (*http.Client, error) {

	basicAuth := func(req *http.Request) (*url.URL, error) {
		req.SetBasicAuth(ba.Username, ba.Password)
		return req.URL, nil
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: basicAuth,
		},
	}, nil
}
