package testdata

import (
	"fmt"
	"github.com/flashmob/go-guerrilla"
	glog "github.com/flashmob/go-guerrilla/log"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

func BirthDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func LocalDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func StartSMTPServer() guerrilla.Daemon {
	cfg := &guerrilla.AppConfig{
		LogFile:      glog.OutputOff.String(),
		AllowedHosts: []string{"test"},
	}

	d := guerrilla.Daemon{Config: cfg}

	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
	}

	return d
}

func TempFile(content string, dir string) string {
	return TempFileWithName(content, dir, "tmp-file")
}

func TempFileWithName(content string, dir string, filename string) string {
	byteContent := []byte(content)
	tmpfn := filepath.Join(dir, filename)
	if err := ioutil.WriteFile(tmpfn, byteContent, 0666); err != nil {
		log.Fatal(err)
	}
	return tmpfn
}

func TempDir() string {
	dir, err := ioutil.TempDir("", "tmp-dir")
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func JsonOauthConfig(clientId string) string {
	return "{\"installed\": " +
		"{\"client_id\": \"" + clientId + "\"," +
		"\"redirect_uris\": [\"http://uri\"]}}"
}

func Oauth2Config(clientID string) *oauth2.Config {
	return &oauth2.Config{ClientID: clientID, ClientSecret: "", Endpoint: oauth2.Endpoint{AuthURL: "", TokenURL: ""}, RedirectURL: "http://uri", Scopes: []string{""}}
}
