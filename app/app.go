package app

import (
	"github.com/gorilla/mux"
	"github.com/srodrigo/payments/router"
)

type App struct {
	Router *mux.Router
}

func CreateApp() App {
	return App{Router: router.NewRouter()}
}
