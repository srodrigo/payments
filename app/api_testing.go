package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func assertEmptyResponseBody(t *testing.T, body *bytes.Buffer) {
	assert.Equal(t, 0, len(body.Bytes()))
}

func readTestFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("../test_data/%s", filename))
}

func (app *App) makeGetRequest(path string) *httptest.ResponseRecorder {
	return app.makeRequestWithBody("GET", path, nil)
}

func (app *App) makePostRequest(path string, payload []byte) *httptest.ResponseRecorder {
	return app.makeRequestWithBody("POST", path, payload)
}

func (app *App) makePutRequest(path string, payload []byte) *httptest.ResponseRecorder {
	return app.makeRequestWithBody("PUT", path, payload)
}

func (app *App) makeDeleteRequest(path string) *httptest.ResponseRecorder {
	return app.makeRequestWithBody("DELETE", path, nil)
}

func (app *App) makeRequestWithBody(method, path string, payload []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(response, req)

	return response
}
