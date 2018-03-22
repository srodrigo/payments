FROM golang:1.10.0-alpine

RUN apk update && apk add git

WORKDIR /go/src/github.com/srodrigo/payments

ADD ./install-go-deps.sh /go/src/github.com/srodrigo/payments
RUN ash ./install-go-deps.sh

ADD . /go/src/github.com/srodrigo/payments

RUN go install

EXPOSE 8000

CMD /go/bin/payments
