package app

import (
	"fmt"
	"github.com/srodrigo/payments/payments"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const BASE_URL = "/v1/payments"

func TestCreateApp(t *testing.T) {
	paymentsRepository := payments.PaymentsRepository{}
	app := CreateApp(&paymentsRepository)

	assert.NotNil(t, app.Router)
}

func TestCreatesPayment(t *testing.T) {
	paymentsRepository := paymentsRepositoryWithIds("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")
	app := CreateApp(paymentsRepository)

	response := app.createPayment("create-payment-1_request.json")

	assertResponseCode(t, http.StatusCreated, response.Code)
	assertResponseBody(t, "payment-1_response.json", response.Body)
}

func TestUpdatesPayment(t *testing.T) {
	id := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
	paymentsRepository := paymentsRepositoryWithIds(id)
	app := CreateApp(paymentsRepository)
	app.createPayment("create-payment-1_request.json")

	// Updates the "amount" field
	response := app.updatePayment(id, "update-payment-1_request.json")

	assertResponseCode(t, http.StatusOK, response.Code)
	assertResponseBody(t, "updated-payment-1_response.json", response.Body)
}

func TestGetsPayment(t *testing.T) {
	paymentsRepository := paymentsRepositoryWithIds("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")
	app := CreateApp(paymentsRepository)
	app.createPayment("create-payment-1_request.json")

	response := app.getPayment("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")

	assertResponseCode(t, http.StatusOK, response.Code)
	assertResponseBody(t, "payment-1_response.json", response.Body)
}

func TestGetAllPayments(t *testing.T) {
	paymentsRepository := paymentsRepositoryWithIds(
		"4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
		"216d4da9-e59a-4cc6-8df3-3da6e7580b77",
	)
	app := CreateApp(paymentsRepository)
	app.createPayment("create-payment-1_request.json")
	app.createPayment("create-payment-2_request.json")

	response := app.getAllPayments()

	assertResponseCode(t, http.StatusOK, response.Code)
	assertResponseBody(t, "all-payments_response.json", response.Body)
}

func TestDeletesPayment(t *testing.T) {
	id := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
	paymentsRepository := paymentsRepositoryWithIds(id)
	app := CreateApp(paymentsRepository)
	app.createPayment("create-payment-1_request.json")

	response := app.deletePayment(id)

	assertResponseCode(t, http.StatusNoContent, response.Code)
	assertEmptyResponseBody(t, response.Body)
}

func TestBadRequestWhenCreatePaymentBodyIsInvalid(t *testing.T) {
	paymentsRepository := paymentsRepositoryWithIds("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")
	app := CreateApp(paymentsRepository)

	response := app.createPayment("create-payment-malformed-body_request.json")

	assertResponseCode(t, http.StatusBadRequest, response.Code)
	assertEmptyResponseBody(t, response.Body)
}

func TestBadRequestWhenUpdatePaymentBodyIsInvalid(t *testing.T) {
	id := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
	paymentsRepository := paymentsRepositoryWithIds(id)
	app := CreateApp(paymentsRepository)

	response := app.updatePayment(id, "create-payment-malformed-body_request.json")

	assertResponseCode(t, http.StatusBadRequest, response.Code)
	assertEmptyResponseBody(t, response.Body)
}

func TestNotFoundWhenUpdatePaymentIdDoesNotExist(t *testing.T) {
	id := "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"
	paymentsRepository := paymentsRepositoryWithIds(id)
	app := CreateApp(paymentsRepository)

	response := app.updatePayment(id, "update-payment-1_request.json")

	assertResponseCode(t, http.StatusNotFound, response.Code)
	assertEmptyResponseBody(t, response.Body)
}

func TestNotFoundWhenDeletePaymentIdDoesNotExist(t *testing.T) {
	paymentsRepository := payments.NewPaymentsRepository()
	app := CreateApp(paymentsRepository)

	response := app.deletePayment("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")

	assertResponseCode(t, http.StatusNotFound, response.Code)
	assertEmptyResponseBody(t, response.Body)
}

func TestNotFoundWhenPaymentDoesNotExist(t *testing.T) {
	paymentsRepository := payments.NewPaymentsRepository()
	app := CreateApp(paymentsRepository)

	response := app.getPayment("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")

	assertResponseCode(t, http.StatusNotFound, response.Code)
	assertEmptyResponseBody(t, response.Body)
}

func (app *App) createPayment(filename string) *httptest.ResponseRecorder {
	payload, err := readTestFile(filename)
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}

	return app.makePostRequest(BASE_URL, payload)
}

func (app *App) updatePayment(id, filename string) *httptest.ResponseRecorder {
	payload, err := readTestFile(filename)
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}

	return app.makePutRequest(fmt.Sprintf("%s/%s", BASE_URL, id), payload)
}

func (app *App) getPayment(id string) *httptest.ResponseRecorder {
	return app.makeGetRequest(fmt.Sprintf("%s/%s", BASE_URL, id))
}

func (app *App) getAllPayments() *httptest.ResponseRecorder {
	return app.makeGetRequest(BASE_URL)
}

func (app *App) deletePayment(id string) *httptest.ResponseRecorder {
	return app.makeDeleteRequest(fmt.Sprintf("%s/%s", BASE_URL, id))
}

type TestUUID struct {
	Ids []string
}

func (uuid *TestUUID) GetNextUUID() string {
	nextUuid := uuid.Ids[0]
	uuid.Ids = uuid.Ids[1:]
	return nextUuid
}

func paymentsRepositoryWithIds(ids ...string) *payments.PaymentsRepository {
	return &payments.PaymentsRepository{
		Uuid: &TestUUID{Ids: ids},
	}
}
