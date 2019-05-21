package request

import (
	"github.com/s-petit/birthday-pal/app/contact"
)

//ContactsProvider calls a HTTP Contact API and returns a slice of Contacts regardless of protocol or authentication
type ContactsProvider interface {
	GetContacts() ([]contact.Contact, error)
}
