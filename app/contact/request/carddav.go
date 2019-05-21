package request

import (
	"errors"
	"github.com/s-petit/birthday-pal/app/contact"
	"github.com/s-petit/birthday-pal/app/contact/auth"
	"github.com/s-petit/birthday-pal/app/contact/request/vcard"
	"io/ioutil"
	"log"
	"strconv"
)

//CardDavContactsProvider represents a provider which return contacts via CardDav protocol
type CardDavContactsProvider struct {
	AuthClient auth.AuthenticationClient
	URL        string
}

//GetContacts returns contacts via a CardDav HTTP API call
func (carddav CardDavContactsProvider) GetContacts() ([]contact.Contact, error) {
	vcards, err := carddav.call()
	if err != nil {
		return []contact.Contact{}, err
	}
	return vcard.ParseContacts(vcards)
}

func (carddav CardDavContactsProvider) call() (string, error) {
	client, err := carddav.AuthClient.Client()
	if err != nil {
		return "", err
	}

	resp, err := client.Get(carddav.URL)
	if err != nil {
		log.Printf("Carddav HTTP GET - Something went wrong with the URL: %s ", carddav.URL)
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
