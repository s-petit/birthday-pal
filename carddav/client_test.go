package carddav

import (
	"testing"
	"fmt"
)

/*func Test_call_cardDav(t *testing.T) {

	payload := RetrieveContacts("lol", "lol", "lol")

	fmt.Println(payload)

	vcards := getCards(payload)
	assert.Equal(t, 142, len(vcards))
	assert.Equal(t, "19831028", vcards[0].BirthDay)
	assert.Equal(t, "19860425", vcards[1].BirthDay)
}*/

func Test_lol(t *testing.T) {
	client := ContactClient{"https://carddav.fastmail.com/dav/addressbooks/user/spetit@enjoycode.fr/Default", "spetit@enjoycode.fr", "9jx4tgvabryracu3"}

	s, _ := client.Get()

	fmt.Println(s)
}
