package request

import (
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/vcard"
)

//CardDavContactsProvider represents a provider which return contacts via CardDav protocol
type CardDavContactsProvider struct {
	Client auth.Client
	URL    string
}

//Get returns contacts via a CardDav API call
func (carddav CardDavContactsProvider) Get() ([]contact.Contact, error) {
	vcards, err := carddav.Client.Get(carddav.URL)
	if err != nil {
		return []contact.Contact{}, err
	}
	return vcard.ParseContacts(vcards)
}
