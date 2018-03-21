package app

import (
	"github.com/gorilla/mux"
	"github.com/srodrigo/payments/payments"
	"github.com/srodrigo/payments/router"
)

type App struct {
	Router *mux.Router
}

func CreateApp(paymentsRepository *payments.PaymentsRepository) App {
	return App{Router: router.NewRouter()}
}
