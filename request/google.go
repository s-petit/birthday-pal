package request

import (
	"fmt"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"google.golang.org/api/gensupport"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

//GoogleContactsProvider represents a provider which return contacts via Google People API
type GoogleContactsProvider struct {
	Client auth.Client
	URL    string
}

//Get returns contacts via a Google People API call
func (gp GoogleContactsProvider) GetContacts() ([]contact.Contact, error) {
	//TODO SPE client.Client... bof
	clt, err := gp.Client.Client()
	if err != nil {
		return []contact.Contact{}, err
	}

	res, err := clt.Get(gp.URL)
	//res, err := gp.Client.Get(gp.URL)
	if err != nil {
		return []contact.Contact{}, err
	}

	ret := &people.ListConnectionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}

	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}

	for _, i := range ret.Connections {
		for _, b := range i.Birthdays {
			fmt.Printf("%s: %d\n", i.Names[0].DisplayName, b.Date.Month)
		}
	}

	return []contact.Contact{}, nil
}
