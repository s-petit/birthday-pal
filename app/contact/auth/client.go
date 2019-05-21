package auth

import "net/http"

//AuthenticationClient can do a HTTP GET with required authentication
type AuthenticationClient interface {
	Client() (*http.Client, error)
}
