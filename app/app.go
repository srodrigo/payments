package app

import (
	"github.com/srodrigo/payments/payments"
	"github.com/srodrigo/payments/router"
)

type App struct {
	Router *router.Router
}

func CreateApp(paymentsRepository *payments.PaymentsRepository) App {
	return App{Router: router.NewRouter(paymentsRepository)}
}
