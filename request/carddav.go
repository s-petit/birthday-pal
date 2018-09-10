package request

import (
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/vcard"
)

//GoogleRequest represents a Google HTTP Request with Basic Auth
type cardDavContactsProvider struct {
	client auth.Client
	URL    string
}

func (gp cardDavContactsProvider) Get() ([]contact.Contact, error) {
	vcards, err := gp.client.Get(gp.URL)
	if err != nil {
		return []contact.Contact{}, err
	}
	return vcard.ParseContacts(vcards)
}
