test:
	go test -v ./...

test-cover:
	go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' ./... | xargs -L 1 sh -c
	gover
	goveralls -coverprofile=gover.coverprofile -service=travis-ci

check: lint vet fmtcheck ineffassign readmecheck

lint:
	golint -set_exit_status ./...

vet:
	go vet

doc:
	autoreadme -f

fmtcheck:
	@ export output="$$(gofmt -s -d .)"; \
		[ -n "$${output}" ] && echo "$${output}" && export status=1; \
		exit $${status:-0}

ineffassign:
	ineffassign .

readmecheck:
	sed '$ d' README.md > README.original.md
	autoreadme -f
	sed '$ d' README.md > README.generated.md
	diff README.generated.md README.original.md

setup:
	go get github.com/gordonklaus/ineffassign
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	go get github.com/modocache/gover
	go get github.com/divan/autoreadme
	go get -t -u ./...

setup-prod:
	go get github.com/s-petit/birthday-pal

build-linux:
	GOARCH=amd64 GOOS=linux go build -o bin/birthday-pal-linux-amd64

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o bin/birthday-pal-mac-amd64

build-windows:
	GOARCH=amd64 GOOS=windows go build -o bin/birthday-pal-windows-amd64.exe

release: build-linux build-mac build-windows
