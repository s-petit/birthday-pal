package http

import (
	"google.golang.org/api/people/v1"
	"log"
	"fmt"
	"github.com/s-petit/birthday-pal/contact"
	"net/http"
	"time"
)


//Request holds methods necessary for requesting cardDAV HTTP servers.
type ContactsProvider interface {
	Get(client *http.Client) ([]contact.Contact, error)
}

//GoogleRequest represents a Google HTTP Request with Basic Auth
type GoogleProvider struct {
	//TODO add client as a field ?
	request      *people.PeopleConnectionsListCall
}

func (gp GoogleProvider) Get(client *http.Client) ([]contact.Contact, error) {

	response := gp.call(client)

	return ParseContacts(response)

}

func  (gp GoogleProvider) call(client *http.Client) (*people.ListConnectionsResponse) {
	people, err := people.New(client)
	if err != nil {
		log.Fatal("could not create the google people service")
	}
	fmt.Println(people)

	response, err := gp.request.Do()
	if err != nil {
		log.Fatal(err)
	}
	return response
}

//TODO interface ? ou mutualiser ce parse avec l'autre parse ?
//ParseContacts parses a cardDav payload to a Contact struct.
func ParseContacts(response *people.ListConnectionsResponse) ([]contact.Contact, error) {

	var contacts []contact.Contact

	for _, connection := range response.Connections {

		for _, b := range connection.Birthdays {
			//fmt.Println(b.Text)
			contacts = append(contacts, contact.Contact{Name: connection.Names[0].DisplayName, BirthDate: time.Date(int(b.Date.Year), time.Month(b.Date.Month), int(b.Date.Day), 0, 0, 0, 0, time.UTC)})
			fmt.Printf("%s: %d\n", connection.Names[0].DisplayName, )
		}
	}

	return contacts, nil
}
