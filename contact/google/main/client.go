package main

import (
	"log"
	"fmt"
	"google.golang.org/api/calendar/v3"
	"time"
	"context"
	"net/http"
	"strconv"
	"io/ioutil"
	"google.golang.org/api/people/v1"
	"github.com/s-petit/birthday-pal/contact/google"
)

func main() {
	// Initialize authentication
	auth := new(google.Authentication)

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

	//https://people.googleapis.com/v1/people/me/connections?requestMask.includeField=person.names%2Cperson.birthdays

	//manualRequest(client)

	googleApiRequest(client)

}

func googleApiRequest(client *http.Client) {
	people, err := people.New(client)
	if err != nil {
		log.Fatal("could not create the google calendar service")
	}
	fmt.Println(people)
	response, err := people.People.Connections.List("people/me").PageSize(500).RequestMaskIncludeField("person.names,person.birthdays").Do()
	if err != nil {
		log.Fatal(err)
	}
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

func manualRequest(client *http.Client) {
	req, err := http.NewRequest("GET", "https://people.googleapis.com/v1/people/me/connections?requestMask.includeField=person.names%2Cperson.birthdays", nil)
	if err != nil {
		log.Fatal(err)
	}
	//client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("a unexpected error occurred during connexion to Google People API server - http code is " + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func cal() {
	// Initialize authentication
	auth := new(google.Authentication)

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

	// Create the google calendar service
	gcal, err := calendar.New(client)
	if err != nil {
		log.Fatal("could not create the google calendar service")
	}

	// Create the time to get events from.
	now := time.Now().Format(time.RFC3339)

	// Get the events list from the calendar service.
	events, err := gcal.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(now).MaxResults(10).OrderBy("startTime").Do()

	if err != nil {
		log.Fatal("unable to retrieve upcoming events: %v", err)
	}

	// Loop over the events and print them out.
	for _, i := range events.Items {
		var when string
		// If the DateTime is an empty string,
		// the event is an all day event
		if i.Start.DateTime != "" {
			when = i.Start.DateTime
		} else {
			when = i.Start.Date
		}
		fmt.Printf("%s (%s)\n", i.Summary, when)
	}
}

