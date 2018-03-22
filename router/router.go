package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/srodrigo/payments/payments"
	"io/ioutil"
	"log"
	"net/http"
)

type Router struct {
	Router *mux.Router
}

type PaymentPayload struct {
	Id             string            `json:"id"`
	Type           string            `json:"type"`
	Version        int               `json:"version"`
	OrganisationId string            `json:"organisation_id"`
	Payment        *payments.Payment `json:"attributes"`
}

type PaymentsListPayload struct {
	Data  []*PaymentPayload `json:"data"`
	Links LinksPayload      `json:"links"`
}

type LinksPayload struct {
	Self string `json:"self"`
}

func NewRouter(paymentsRepository *payments.PaymentsRepository) *Router {
	paymentsService := payments.NewPaymentsService(paymentsRepository)
	const BASE_URL = "/v1/payments"

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc(BASE_URL, GetAllPaymentsHandler(paymentsService)).Methods("GET")
	muxRouter.HandleFunc(BASE_URL+"/{id}", GetPaymentHandler(paymentsService)).Methods("GET")
	muxRouter.HandleFunc(BASE_URL, CreatePaymentHandler(paymentsService)).Methods("POST")
	muxRouter.HandleFunc(BASE_URL+"/{id}", UpdatePaymentHandler(paymentsService)).Methods("PUT")
	muxRouter.HandleFunc(BASE_URL+"/{id}", DeletePaymentHandler(paymentsService)).Methods("DELETE")

	return &Router{
		Router: muxRouter,
	}
}

func GetAllPaymentsHandler(paymentsService *payments.PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newPayment := paymentsService.GetAllPayments()

		url := fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)
		payload, _ := createAllPaymentsPayload(newPayment, url)

		writeJsonResponse(w, http.StatusOK, payload)
	}
}

func GetPaymentHandler(paymentsService *payments.PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		newPayment, err := paymentsService.GetPaymentById(vars["id"])
		if err != nil {
			log.Println("Error getting payment:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		payload, _ := marshalPayment(newPayment)

		writeJsonResponse(w, http.StatusOK, payload)
	}
}

func CreatePaymentHandler(paymentsService *payments.PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Handle error
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		var payment payments.Payment
		err := json.Unmarshal(b, &payment)
		if err != nil {
			log.Println("Error parsing body:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newPayment := paymentsService.CreatePayment(&payment)

		payload, _ := marshalPayment(newPayment)

		writeJsonResponse(w, http.StatusCreated, payload)
	}
}

func UpdatePaymentHandler(paymentsService *payments.PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Handle error
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		var payment payments.Payment
		err := json.Unmarshal(b, &payment)
		if err != nil {
			log.Println("Error parsing body:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		updatedPayment, err := paymentsService.UpdatePayment(vars["id"], &payment)
		if err != nil {
			log.Println("Error updating payment:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		payload, _ := marshalPayment(updatedPayment)

		writeJsonResponse(w, http.StatusOK, payload)
	}
}

func DeletePaymentHandler(paymentsService *payments.PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Handle error
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		var payment payments.Payment
		// TODO: Handle error
		json.Unmarshal(b, &payment)

		vars := mux.Vars(r)
		err := paymentsService.DeletePayment(vars["id"])
		if err != nil {
			log.Println("Error updating payment:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func writeJsonResponse(w http.ResponseWriter, statusCode int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(payload)
}

func marshalPayment(payment *payments.Payment) ([]byte, error) {
	return json.Marshal(createPaymentPayload(payment))
}

func marshalPaymentPayload(paymentPayload *PaymentPayload) ([]byte, error) {
	return json.Marshal(*paymentPayload)
}

func createPaymentPayload(payment *payments.Payment) *PaymentPayload {
	return &PaymentPayload{
		Id:             payment.Id,
		Type:           "Payment",
		Version:        0,
		OrganisationId: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Payment:        payment,
	}
}

func createAllPaymentsPayload(payments []*payments.Payment, url string) ([]byte, error) {
	paymentsPayload := make([]*PaymentPayload, len(payments))
	for i := 0; i < len(payments); i++ {
		payload := createPaymentPayload(payments[i])
		paymentsPayload[i] = payload
	}

	return json.Marshal(PaymentsListPayload{
		Data:  paymentsPayload,
		Links: LinksPayload{Self: url},
	})
}
