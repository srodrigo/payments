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

func TestCreateApp(t *testing.T) {
	paymentsRepository := payments.PaymentsRepository{}
	app := CreateApp(&paymentsRepository)

	assert.NotNil(t, app.Router)
}

func TestCreatesPayment(t *testing.T) {
	paymentsRepository := payments.PaymentsRepository{}
	app := CreateApp(&paymentsRepository)

	response := app.createPayment("create-payment-1_request.json")

	assert.Equal(t, http.StatusCreated, response.Code)

	var responseBody map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	expectedJson, err := readTestFile("payment-1_response.json")
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}
	var expected map[string]interface{}
	json.Unmarshal(expectedJson, &expected)
	assert.Equal(t, expected, responseBody)
}

func TestGetsPayment(t *testing.T) {
	paymentsRepository := payments.PaymentsRepository{}
	app := CreateApp(&paymentsRepository)
	app.createPayment("create-payment-1_request.json")

	response := app.getPayment()

	assert.Equal(t, http.StatusOK, response.Code)

	var responseBody map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	expectedJson, err := readTestFile("payment-1_response.json")
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}
	var expected map[string]interface{}
	json.Unmarshal(expectedJson, &expected)
	assert.Equal(t, expected, responseBody)
}

func TestGetAllPayments(t *testing.T) {
	paymentsRepository := payments.PaymentsRepository{}
	app := CreateApp(&paymentsRepository)
	app.createPayment("create-payment-1_request.json")
	app.createPayment("create-payment-2_request.json")

	response := app.getAllPayments()

	assert.Equal(t, http.StatusOK, response.Code)

	var responseBody map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	expectedJson, err := readTestFile("all-payments_response.json")
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}
	var expected map[string]interface{}
	json.Unmarshal(expectedJson, &expected)
	assert.Equal(t, expected, responseBody)
}

func (app *App) createPayment(filename string) *httptest.ResponseRecorder {
	payload, err := readTestFile(filename)
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(response, req)

	return response
}

func (app *App) getPayment() *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", nil)
	response := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(response, req)

	return response
}

func (app *App) getAllPayments() *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/payments", nil)
	response := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(response, req)

	return response
}

func readTestFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("../test_data/%s", filename))
}
