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

	assert.Equal(t, responseBody["id"], "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43")
}

func (app *App) createPayment() *httptest.ResponseRecorder {
	payload, err := readTestFile("create-payment-1.json")
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
