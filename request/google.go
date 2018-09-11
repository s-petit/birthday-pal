package request

import (
	"fmt"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/contact"
	"google.golang.org/api/gensupport"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

//GoogleContactsProvider represents a provider which return contacts via Google People API
type GoogleContactsProvider struct {
	Client auth.Client
	URL    string
}

//Get returns contacts via a Google People API call
func (gp GoogleContactsProvider) Get() ([]contact.Contact, error) {
	clt, err := gp.Client.Clt(gp.URL)
	res, err := clt.Get(gp.URL)
	//res, err := gp.Client.Get(gp.URL)
	if err != nil {
		return []contact.Contact{}, err
	}

	ret := &people.ListConnectionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}

	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}

	for _, i := range ret.Connections {
		/*		for _, n := range i.Names {
				fmt.Println(n.DisplayName)
				//fmt.Printf("%s: %s\n", n.DisplayName, i.Birthdays[0].Text)
			}*/
		for _, b := range i.Birthdays {
			//fmt.Println(b.Text)

			fmt.Printf("%s: %d\n", i.Names[0].DisplayName, b.Date.Month)
		}
		//var when string
		//fmt.Printf("%s: %s (%s)\n", i.Names, i.Birthdays, when)
	}

	//googleapi.ServerResponse{}
	//people.ListConnectionsResponse{}
	return []contact.Contact{}, nil
}

/*
func (gp GoogleContactsProvider) call() ([]contact.Contact, error) {
	people, err := people.New(gp.client.)
	if err != nil {
		log.Fatal("could not create the google people service")
	}
	fmt.Println(people)

	req := people.People.Connections.List("people/me").PageSize(500).RequestMaskIncludeField("person.names,person.birthdays")
	res, err := req.Do()
	if err != nil {
		log.Fatal(err)
	}

	ret := &people.ListConnectionsResponse{
		googleapi.ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}


	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil

	return response
}*/
