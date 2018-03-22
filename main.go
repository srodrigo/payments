package main

import (
	"github.com/srodrigo/payments/app"
	"github.com/srodrigo/payments/payments"
)

func main() {
	paymentsRepository := payments.PaymentsRepository{}
	server := app.CreateApp(&paymentsRepository)
	server.Run()
}
