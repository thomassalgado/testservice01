package main

import "testing"
import "net/http"
import "net/http/httptest"
import "github.com/stretchr/testify/assert"


var errorMessages = []struct {
	begin  string
	end string
  message string
}{
	{"A", "100", "invalid begin value"},
  {"0", "A", "invalid end value"},
  {"-10", "100", "parameters must be unsigned and greater than 0"},
  {"90", "10", "begin must be lower than end"},
  {"105", "110", "parameters must be lower than 100"},
}

func TestInvalidBegin(t *testing.T) {
  for _, errorMessage := range errorMessages {
    message, err := processNumbers(errorMessage.begin, errorMessage.end)
    assert.Equalf(t, message, "", "message must be empty")
    assert.NotNilf(t, err, "error must be not nil")
    assert.Equalf(t, err.Error(), errorMessage.message, "error message must be %s", "errorMessage.message")
	}
}

func TestSuccess(t *testing.T) {
    begin := "1"
    end := "15"
    expectedMessage := "1 2 Pé 4 Do Pé 7 8 Pé Do 11 Pé 13 14 PéDo"
    message, err := processNumbers(begin, end)
    assert.Equalf(t, message, expectedMessage, "message must be empty %s", expectedMessage)
    assert.Nilf(t, err, "error must be nil")
}

func TestSuccessRequest(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/numbers", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestErrorRequest(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/numbers?begin=-1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
