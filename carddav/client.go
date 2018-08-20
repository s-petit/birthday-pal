package carddav

import (
	"net/http"
	"io/ioutil"
)

func Contacts(url string, username string, password string) string {

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

