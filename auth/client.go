package auth

import "net/http"

//AuthClient can do a HTTP GET with required authentication
type AuthClient interface {
	Client() (*http.Client, error)
}
