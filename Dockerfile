FROM golang:latest

ENV GOROOT /usr/local/go
ENV GOPATH /go
COPY . src/github.com/s-petit/birthday-pal/
WORKDIR src/github.com/s-petit/birthday-pal
RUN make setup-prod
