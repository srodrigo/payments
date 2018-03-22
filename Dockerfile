FROM golang:1.10.0-alpine

RUN apk update && apk add git

RUN go get github.com/stretchr/testify/assert \
  github.com/gorilla/mux \
  github.com/satori/go.uuid

WORKDIR /go/src/github.com/srodrigo/payments
ADD . /go/src/github.com/srodrigo/payments

RUN go install

EXPOSE 8000

CMD /go/bin/payments
