package google

import (
	"github.com/s-petit/birthday-pal/contact"
	"google.golang.org/api/people/v1"
	"fmt"
	"time"
)

//TODO interface ? ou mutualiser ce parse avec l'autre parse ?
//ParseContacts parses a cardDav payload to a Contact struct.
func ParseContacts(response people.ListConnectionsResponse) ([]contact.Contact, error) {

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
