package google

import (
	"log"
	"net/http"
	"fmt"
	"google.golang.org/api/calendar/v3"
	"time"
	"context"
)

func main() {
	// Initialize authentication
	auth := new(Authentication)

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

	req, err := http.NewRequest("PROPFIND", "https://www.googleapis.com/.well-known/carddav", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(resp.StatusCode)
}

func cal() {
	// Initialize authentication
	auth := new(Authentication)

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

