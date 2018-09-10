package auth

//Client can do a HTTP GET with required authentication
type Client interface {
	Get(url string) (string, error)
}
