package main

import (
	"context"
	"fmt"
	bpalgoogle "github.com/s-petit/birthday-pal/http/google"
	"google.golang.org/api/people/v1"
	"log"
	"net/http"
)

func main() {
	client := oauthClient(people.ContactsReadonlyScope)
	googleApiRequest(client)
}

func oauthClient(scope string) *http.Client {
	// Initialize authentication
	auth := new(bpalgoogle.Authentication)
	auth.Scope = scope

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

//json.NewDecoder(res.Body).Decode(target)

func googleApiRequest(client *http.Client)  {
	people, err := people.New(client)
	if err != nil {
		log.Fatal("could not create the google calendar service")
	}
	fmt.Println(people)
	response, err := people.People.Connections.List("people/me").PageSize(500).RequestMaskIncludeField("person.names,person.birthdays").Do()
	if err != nil {
		log.Fatal(err)
	}

	Connections := response.Connections
	fmt.Println(response.TotalPeople)
	fmt.Println(len(response.Connections))

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
