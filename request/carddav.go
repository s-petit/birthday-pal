package request

import (
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/vcard"
)

//GoogleRequest represents a Google HTTP Request with Basic Auth
type CardDavContactsProvider struct {
	Client auth.Client
	URL    string
}

func (carddav CardDavContactsProvider) Get() ([]contact.Contact, error) {
	vcards, err := carddav.Client.Get(carddav.URL)
	if err != nil {
		return []contact.Contact{}, err
	}
	return vcard.ParseContacts(vcards)
}
