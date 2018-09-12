package request

import (
	"github.com/s-petit/birthday-pal/contact"
)

//ContactsProvider returns a slice of Contacts regardless of protocol or authentication
type ContactsProvider interface {
	GetContacts() ([]contact.Contact, error)
}

//NewContactsProvider create the right instance of ContactsProvider depending on the url
/*func NewContactsProvider(URL string, AuthClient auth.AuthClient) ContactsProvider {

	googleAPI := regexp.MustCompile("https://(\\w+).googleapis.com/(\\w+)")

	if googleAPI.MatchString(URL) {
		return GoogleContactsProvider{AuthClient: AuthClient, URL: URL}
	}

	return CardDavContactsProvider{AuthClient: AuthClient, URL: URL}
}*/
