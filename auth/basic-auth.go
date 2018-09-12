package auth

import (
	"net/http"
)

//TODO revoir la godoc
//TODO SPE revoir la visibilite de la plupart des fields et methods

//BasicAuth provides
type BasicAuth struct {
	Username string
	Password string
}

func (ba BasicAuth) Client() (*http.Client, error) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(ba.Username, ba.Password)

	return &http.Client{}, err
}

//Get invokes a HTTP Get with BasicAuth and handles errors
/*
*/
