package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/srodrigo/payments/payments"
	"io/ioutil"
	"net/http"
)

type Router struct {
	Router *mux.Router
}

func NewRouter(paymentsRepository *payments.PaymentsRepository) *Router {
	paymentsService := payments.NewPaymentsService(paymentsRepository)

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/payments", CreatePaymentHandler(paymentsService)).Methods("POST")

	return &Router{
		Router: muxRouter,
	}
}

func CreatePaymentHandler(paymentsService *payments.PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Handle error
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		var payment payments.Payment
		// TODO: Handle error
		json.Unmarshal(b, &payment)

		newPayment := paymentsService.CreatePayment(&payment)

		// TODO: Handle error
		payload, _ := createPaymentPayload(newPayment)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(payload)
	}
}

func createPaymentPayload(payment *payments.Payment) ([]byte, error) {
	return json.Marshal(Payload{
		Id:             payment.Id,
		Type:           "Payment",
		Version:        0,
		OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Attributes:     payment,
	})
}
