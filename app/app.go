package app

import (
	"fmt"
	"github.com/srodrigo/payments/payments"
	"github.com/srodrigo/payments/router"
	"log"
	"net/http"
)

type App struct {
	Router *router.Router
}

func CreateApp(paymentsRepository *payments.PaymentsRepository) App {
	return App{Router: router.NewRouter(paymentsRepository)}
}

func (app *App) Run() {
	port := 8000
	log.Println(fmt.Sprintf("Listening on port %d", port))

	http.ListenAndServe(fmt.Sprintf(":%d", port), app.Router.Router)
}
