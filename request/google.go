package request

import (
	"github.com/s-petit/birthday-pal/contact"
	"google.golang.org/api/gensupport"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
	"time"
	"github.com/s-petit/birthday-pal/auth"
)

//GoogleContactsProvider represents a provider which return contacts via Google People API
type GoogleContactsProvider struct {
	Client auth.AuthClient
	URL    string
}

//Get returns contacts via a Google People API call
func (gp GoogleContactsProvider) GetContacts() ([]contact.Contact, error) {
	clt, err := gp.Client.Client()
	if err != nil {
		return []contact.Contact{}, err
	}

	res, err := clt.Get(gp.URL)
	if err != nil {
		return []contact.Contact{}, err
	}

	connections := &people.ListConnectionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}

	target := &connections
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}

	return parseContacts(connections)
}

func parseContacts(connections *people.ListConnectionsResponse) ([]contact.Contact, error) {

	var contacts []contact.Contact

	for _, connection := range connections.Connections {
		for _, b := range connection.Birthdays {
			contacts = append(contacts, contact.Contact{Name: connection.Names[0].DisplayName, BirthDate: time.Date(int(b.Date.Year), time.Month(b.Date.Month), int(b.Date.Day), 0, 0, 0, 0, time.UTC)})
		}
	}

	return contacts, nil
}
