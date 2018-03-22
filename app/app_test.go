package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/srodrigo/payments/payments"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
	paymentsRepository := payments.PaymentsRepository{}
	app := CreateApp(&paymentsRepository)

	response := app.createPayment("create-payment-1_request.json")

	assertResponseCode(t, http.StatusCreated, response.Code)
	assertResponseBody(t, "payment-1_response.json", response.Body)
}

func TestGetsPayment(t *testing.T) {
	paymentsRepository := payments.PaymentsRepository{}
	app := CreateApp(&paymentsRepository)
	app.createPayment("create-payment-1_request.json")

	response := app.getPayment("4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")

	assertResponseCode(t, http.StatusOK, response.Code)
	assertResponseBody(t, "payment-1_response.json", response.Body)
}

func TestGetAllPayments(t *testing.T) {
	paymentsRepository := payments.PaymentsRepository{}
	app := CreateApp(&paymentsRepository)
	app.createPayment("create-payment-1_request.json")
	app.createPayment("create-payment-2_request.json")

	response := app.getAllPayments()

	assertResponseCode(t, http.StatusOK, response.Code)
	assertResponseBody(t, "all-payments_response.json", response.Body)
}

func (app *App) createPayment(filename string) *httptest.ResponseRecorder {
	payload, err := readTestFile(filename)
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}

	return app.makePostRequest(BASE_URL, payload)
}

func (app *App) getPayment(id string) *httptest.ResponseRecorder {
	return app.makeGetRequest(fmt.Sprintf("%s/%s", BASE_URL, id))
}

func (app *App) getAllPayments() *httptest.ResponseRecorder {
	return app.makeGetRequest(BASE_URL)
}

func assertResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual)
}

func assertResponseBody(t *testing.T, filename string, body *bytes.Buffer) {
	var responseBody map[string]interface{}
	json.Unmarshal(body.Bytes(), &responseBody)

	expectedJson, err := readTestFile(filename)
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}

	var expected map[string]interface{}
	err = json.Unmarshal(expectedJson, &expected)
	if err != nil {
		fmt.Println("Error parsing data")
		fmt.Println(err)
	}

	assert.Equal(t, expected, responseBody)
}

func readTestFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("../test_data/%s", filename))
}

func (app *App) makeGetRequest(path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(response, req)

	return response
}

func (app *App) makePostRequest(path string, payload []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(response, req)

	return response
}
