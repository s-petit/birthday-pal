package main

import (
	"testing"
	"os"
	"github.com/s-petit/birthday-pal/carddav"
	"github.com/golang/mock/gomock"
)

//TODO refacto sur le projet entier : privilegier les pointeurs
//TODO ajouter de la validation sur les args, notamment urls et emails
// https://goreportcard.com/report/github.com/vektra/mockery

func Test_main(t *testing.T) {
	//TODO how to test with mow.cli ?
	os.Args = []string{"", "https://mycarddav/com/contacts", "carddav-user", "carddav-pass", "recipient-email"}
/*	os.Args[1] = "https://mycarddav.com/contacts"
	os.Args[2] = "carddav-user"
	os.Args[3] = "carddav-pass"
	os.Args[4] = "recipient-email"*/

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_carddav.NewMockClient(ctrl)
	client.EXPECT().Get("url", "user", "pass").Return("myContacts")
	main()
}
