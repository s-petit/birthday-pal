test:
	go test -v ./...

setup:
	go get github.com/jawher/mow.cli
	go get github.com/mapaiva/vcard-go
	go get -t -u ./...
