package auth

import "net/http"

//Client can do a HTTP GET with required authentication
type Client interface {
	Client() (*http.Client, error)
}
