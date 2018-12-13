package request

import (
	"fmt"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"google.golang.org/api/gensupport"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
	"log"
	"time"
)

//GoogleContactsProvider represents a provider which return contacts via Google People API
type GoogleContactsProvider struct {
	AuthClient auth.AuthenticationClient
	URL        string
}

//GetContacts returns contacts via a Google People API call
func (gp GoogleContactsProvider) GetContacts() ([]contact.Contact, error) {
	clt, err := gp.AuthClient.Client()
	if err != nil {
		return []contact.Contact{}, err
	}

	res, err := clt.Get(gp.URL)
	if err != nil {
		log.Printf("Carddav HTTP GET - Something went wrong with the URL: %s ", gp.URL)
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
		log.Printf("Failed to decode HTTP Response - Something went wrong with the URL: %s ", gp.URL)
		return nil, err
	}

	return parseContacts(connections), nil
}

func parseContacts(connections *people.ListConnectionsResponse) []contact.Contact {

	var contacts []contact.Contact

	for _, connection := range connections.Connections {
		for _, b := range connection.Birthdays {
			if connection.Names[0] != nil && b.Date != nil {
				contacts = append(contacts, contact.Contact{Name: connection.Names[0].DisplayName, BirthDate: time.Date(int(b.Date.Year), time.Month(b.Date.Month), int(b.Date.Day), 0, 0, 0, 0, time.UTC)})
			} else {
				fmt.Println("This contact is malformed: ", connection.Names[0], b)
			}
		}
	}

	return contacts
}
