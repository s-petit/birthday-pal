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



//json.NewDecoder(res.Body).Decode(target)


