package request

import (
	"errors"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/vcard"
	"io/ioutil"
	"strconv"
)

//CardDavContactsProvider represents a provider which return contacts via CardDav protocol
type CardDavContactsProvider struct {
	Client auth.Client
	URL    string
}

//Get returns contacts via a CardDav API call
func (carddav CardDavContactsProvider) GetContacts() ([]contact.Contact, error) {
	vcards, err := carddav.call()
	if err != nil {
		return []contact.Contact{}, err
	}
	return vcard.ParseContacts(vcards)
}

func (carddav CardDavContactsProvider) call() (string, error) {
	client, err := carddav.Client.Client()
	if err != nil {
		return "", err
	}

	resp, err := client.Get(carddav.URL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("a unexpected error occurred during connexion to Contact API server - http code is " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

