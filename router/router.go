package router

import (
	"github.com/gorilla/mux"
	"github.com/srodrigo/payments/payments"
	"net/http"
)

type Router struct {
	Router             *mux.Router
	PaymentsRepository *payments.PaymentsRepository
}

func NewRouter(paymentsRepository *payments.PaymentsRepository) *Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/payments", CreatePayment).Methods("POST")

	return &Router{
		Router:             muxRouter,
		PaymentsRepository: paymentsRepository,
	}
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
}
