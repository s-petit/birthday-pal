## Feature or refactoring ideas

## Features ideas

- Make bpal work with Google's cardDav servers which use oauth2
- implement a real i18n and templating solution
- send a "digest" reminder for a given period. Example : Here are the birthdays of the week...
- un vrai readme pro



//TODO google Request : trouver une solution pour la page size...
/*func (oa OAuth2) googleApiRequest() {
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
		}
		for _, b := range i.Birthdays {
			//fmt.Println(b.Text)

			fmt.Printf("%s: %d\n", i.Names[0].DisplayName, b.Date.Month)
		}
		//var when string
		//fmt.Printf("%s: %s (%s)\n", i.Names, i.Birthdays, when)
	}
}
*/



//TODO google Request : trouver une solution pour la page size...
/*func (oa OAuth2) googleApiRequest() {
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
		}
		for _, b := range i.Birthdays {
			//fmt.Println(b.Text)

			fmt.Printf("%s: %d\n", i.Names[0].DisplayName, b.Date.Month)
		}
		//var when string
		//fmt.Printf("%s: %s (%s)\n", i.Names, i.Birthdays, when)
	}
}
*/


/*func (gp GoogleContactsProvider) call(client *http.Client) *people.ListConnectionsResponse {
	people, err := people.New(client)
	if err != nil {
		log.Fatal("could not create the google people service")
	}
	fmt.Println(people)

	req := people.People.Connections.List("people/me").PageSize(500).RequestMaskIncludeField("person.names,person.birthdays")
	response, err := req.Do()
	if err != nil {
		log.Fatal(err)
	}
	return response
}*/

/*func main() {
	client := oauthClient(people.ContactsReadonlyScope)
	googleApiRequest(client)
}
*/
//json.NewDecoder(res.Body).Decode(target)


//TODO interface ? ou mutualiser ce parse avec l'autre parse ?
//ParseContacts parses a cardDav payload to a Contact struct.
func getContacts(response *people.ListConnectionsResponse) ([]Contact, error) {

	var contacts []Contact

	for _, connection := range response.Connections {
		for _, b := range connection.Birthdays {
			contacts = append(contacts, Contact{Name: connection.Names[0].DisplayName, BirthDate: time.Date(int(b.Date.Year), time.Month(b.Date.Month), int(b.Date.Day), 0, 0, 0, 0, time.UTC)})
		}
	}

	return contacts, nil
}
