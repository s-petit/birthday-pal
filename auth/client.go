package auth

import "net/http"

//Client can do a HTTP GET with required authentication
type Client interface {
	Get(url string) (string, error)
	Clt(url string) (*http.Client, error)
}
