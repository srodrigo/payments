package main

import (
	"github.com/srodrigo/payments/app"
	"github.com/srodrigo/payments/payments"
)

func main() {
	paymentsRepository := payments.NewPaymentsRepository()
	server := app.CreateApp(paymentsRepository)
	server.Run()
}
