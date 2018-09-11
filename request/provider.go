package request

import (
	"github.com/s-petit/birthday-pal/contact"
)

//ContactsProvider returns a slice of Contacts regardless of protocol or authentication
type ContactsProvider interface {
	Get() ([]contact.Contact, error)
}

//NewContactsProvider create the right instance of ContactsProvider depending on the url
/*func NewContactsProvider(URL string, client auth.Client) ContactsProvider {

	googleAPI := regexp.MustCompile("https://(\\w+).googleapis.com/(\\w+)")

	if googleAPI.MatchString(URL) {
		return googleContactsProvider{client: client, URL: URL}
	}

	return CardDavContactsProvider{client: client, URL: URL}
}*/
