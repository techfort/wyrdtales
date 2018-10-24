FROM golang:1.11

MAINTAINER Joe Minichino <joe.minichino@gmail.com>


ADD ./ /go/src/github.com/techfort/wyrdtales

WORKDIR /go/src/github.com/techfort/wyrdtales

RUN go get github.com/canthefason/go-watcher && go install github.com/canthefason/go-watcher/cmd/watcher

CMD go run -race main.go
