package router

import (
	"github.com/gorilla/mux"
	"github.com/srodrigo/payments/payments"
	"net/http"
)

type Router struct {
	Router   *mux.Router
	Payments *payments.Payments
}

func NewRouter(paymentsRepository *payments.PaymentsRepository) *Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/payments", CreatePayment).Methods("POST")

	payments := payments.NewPayments(paymentsRepository)

	return &Router{
		Router:   muxRouter,
		Payments: payments,
	}
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
}
