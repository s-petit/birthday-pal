package http

import (
	"google.golang.org/api/people/v1"
	"log"
	"fmt"
	"github.com/s-petit/birthday-pal/contact"
	"net/http"
	"time"
	"errors"
	"strconv"
	"io/ioutil"
	"github.com/s-petit/birthday-pal/contact/vcard"
)


//Request holds methods necessary for requesting cardDAV HTTP servers.
type ContactsProvider interface {
	Get(client *http.Client) ([]contact.Contact, error)
}

//GoogleRequest represents a Google HTTP Request with Basic Auth
type CardDavContactsProvider struct {
	Request      *http.Request
}

/*
*/
//GoogleRequest represents a Google HTTP Request with Basic Auth
type GoogleContactsProvider struct {
	request      *people.PeopleConnectionsListCall
}

func (gp CardDavContactsProvider) Get(client *http.Client) ([]contact.Contact, error) {
	vcards, err := gp.call(client)
	if err != nil {
		return []contact.Contact{}, err
	}
	return vcard.ParseContacts(vcards)
}


func (gp CardDavContactsProvider) call(client *http.Client) (string, error) {
	resp, err := client.Do(gp.request)
	if err != nil {
	return "", err
	}

	if resp.StatusCode != 200 {
	return "", errors.New("a unexpected error occurred during connexion to CardDAV server - http code is " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	return "", err
	}

	return string(body), nil
}


func (gp GoogleContactsProvider) Get(client *http.Client) ([]contact.Contact, error) {
	response := gp.call(client)
	return getContacts(response)
}

func  (gp GoogleContactsProvider) call(client *http.Client) (*people.ListConnectionsResponse) {
	people, err := people.New(client)
	if err != nil {
		log.Fatal("could not create the google people service")
	}
	fmt.Println(people)

	response, err := gp.request.Do()
	if err != nil {
		log.Fatal(err)
	}
	return response
}

//TODO interface ? ou mutualiser ce parse avec l'autre parse ?
//ParseContacts parses a cardDav payload to a Contact struct.
func getContacts(response *people.ListConnectionsResponse) ([]contact.Contact, error) {

	var contacts []contact.Contact

	for _, connection := range response.Connections {
		for _, b := range connection.Birthdays {
			contacts = append(contacts, contact.Contact{Name: connection.Names[0].DisplayName, BirthDate: time.Date(int(b.Date.Year), time.Month(b.Date.Month), int(b.Date.Day), 0, 0, 0, 0, time.UTC)})
		}
	}

	return contacts, nil
}
