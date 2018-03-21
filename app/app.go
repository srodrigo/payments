package app

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

func CreateApp() App {
	return App{}
}
