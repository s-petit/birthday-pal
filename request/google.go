package request

import (
	"fmt"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
)

//GoogleContactsProvider represents a provider which return contacts via Google People API
type GoogleContactsProvider struct {
	Client auth.Client
	URL    string
}

//Get returns contacts via a Google People API call
func (gp GoogleContactsProvider) Get() ([]contact.Contact, error) {
	response, err := gp.Client.Get(gp.URL)
	if err != nil {
		return []contact.Contact{}, err
	}

	fmt.Println(response)
	//googleapi.ServerResponse{}
	//people.ListConnectionsResponse{}
	return []contact.Contact{}, nil
}
