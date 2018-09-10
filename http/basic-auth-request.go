package http

import (
	"net/http"
	"errors"
	"strconv"
	"io/ioutil"
	"github.com/s-petit/birthday-pal/http/google"
	"log"
	"context"
	"google.golang.org/api/people/v1"
	"fmt"
	"github.com/s-petit/birthday-pal/contact"
)

//TODO revoir la godoc

//Request holds methods necessary for requesting cardDAV HTTP servers.
type AuthProvider interface {
	Get(client *http.Client) ([]contact.Contact, error)
}

//BasicAuth represents a HTTP Request with Basic Auth
type BasicAuth struct {
	Username string
	Password string
}

func (r BasicAuth) Request(URL string) (*http.Request) {
	httpRequest, err := http.NewRequest("GET", URL, nil)
	//TODO log.fatal here ?
	if err != nil {
		log.Fatal(err)
	}
	httpRequest.SetBasicAuth(r.Username, r.Password)

	return httpRequest
}

//Get invokes a HTTP Get with BasicAuth and handles errors
func (r BasicAuth) Get(url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(r.Username, r.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
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



//OAuth2 represents a HTTP Request with OAuth2
type OAuth2 struct {
	scope      string
	auth       google.Authentication
}


func (oa OAuth2) oauthClient() *http.Client {
	// Initialize authentication
/*	auth := new(bpalgoogle.Authentication)
	auth.Scope = scope*/

	//TODO DO IT BETTER
	auth := oa.auth
	auth.Scope = oa.scope

	// Load the configuration from client_secret.json
	config, err := auth.Config()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Load the token from the cache or force authentication
	token, err := auth.Token()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create the API client with a background context.
	ctx := context.Background()
	client := config.Client(ctx, token)

	return client

	//https://people.googleapis.com/v1/people/me/connections?requestMask.includeField=person.names%2Cperson.birthdays
	//googleApiRequest(client)

}

//TODO google request : trouver une solution pour la page size...
func (oa OAuth2) googleApiRequest()  {
	people, err := people.New(oa.oauthClient())
	if err != nil {
		log.Fatal("could not create the google people service")
	}
	fmt.Println(people)

	lol := people.People.Connections.List("people/me").PageSize(500).RequestMaskIncludeField("person.names,person.birthdays")
	response, err := people.People.Connections.List("people/me").PageSize(500).RequestMaskIncludeField("person.names,person.birthdays").Do()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the events and print them out.
	for _, i := range response.Connections {
		/*for _, n := range i.Names {
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
}
