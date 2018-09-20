package auth

import (
	"net/http"
	"net/url"
)

//BasicAuth is used to perform basic authentication
type BasicAuth struct {
	Username string
	Password string
}

//Client returns a HTTP client authenticated with basic Auth
func (ba BasicAuth) Client() (*http.Client, error) {

	basicAuth := func(req *http.Request) (*url.URL, error) {
		req.SetBasicAuth(ba.Username, ba.Password)
		return http.ProxyFromEnvironment(req)
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: basicAuth,
		},
	}, nil
}
