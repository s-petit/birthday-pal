package request

import (
	"fmt"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
)

//GoogleRequest represents a Google HTTP Request with Basic Auth
type googleContactsProvider struct {
	client auth.Client
	URL    string
}

func (gp googleContactsProvider) Get() ([]contact.Contact, error) {
	response, err := gp.client.Get(gp.URL)
	if err != nil {
		return []contact.Contact{}, err
	}

	fmt.Println(response)
	//googleapi.ServerResponse{}
	//people.ListConnectionsResponse{}
	return []contact.Contact{}, nil
}
