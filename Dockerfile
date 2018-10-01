FROM golang:1.10-alpine

ADD . /go/src/github.com/centrifuge/functional-testing
WORKDIR /go/src/github.com/centrifuge/functional-testing

CMD go test ./...