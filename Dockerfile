FROM golang:1.10.0-alpine

WORKDIR /go/src/github.com/srodrigo/payments
ADD . /go/src/github.com/srodrigo/payments

RUN go install

EXPOSE 8000

CMD /go/bin/payments
