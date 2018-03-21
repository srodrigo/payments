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

	response := app.createPayment()

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

func (app *App) createPayment() *httptest.ResponseRecorder {
	payload, err := readTestFile("create-payment-1_request.json")
	if err != nil {
		fmt.Println("Error loading data")
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(response, req)

	return response
}

func readTestFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("../test_data/%s", filename))
}
