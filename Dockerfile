FROM golang:1.10.0-alpine

RUN apk update && apk add git

WORKDIR /go/src/github.com/srodrigo/payments
ADD . /go/src/github.com/srodrigo/payments

RUN ash ./install-go-deps.sh

RUN go install

EXPOSE 8000

CMD /go/bin/payments
