package handlers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/agcalc/gorest/pkg/handlers"
)

func TestHelloWorld(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/", nil)
	handlers.Hello(response, request)
	assert.Equal(t, "Hello World!", response.Body.String())
}

func TestAdd(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/add/3/2", nil)
	request = mux.SetURLVars(request, map[string]string{"a": "3", "b": "2"})
	handlers.Add(response, request)
	assert.Equal(t, "5", response.Body.String())
}

func TestSub(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/sub/3/2", nil)
	request = mux.SetURLVars(request, map[string]string{"a": "3", "b": "2"})
	handlers.Sub(response, request)
	assert.Equal(t, "1", response.Body.String())
}

func TestMult(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/mult/3/2", nil)
	request = mux.SetURLVars(request, map[string]string{"a": "3", "b": "2"})
	handlers.Mult(response, request)
	assert.Equal(t, "6", response.Body.String())
}

func TestDiv(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/div/3/2", nil)
	request = mux.SetURLVars(request, map[string]string{"a": "3", "b": "2"})
	handlers.Div(response, request)
	assert.Equal(t, "1.5", response.Body.String())
}

func TestDivZero(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/div/3/0", nil)
	request = mux.SetURLVars(request, map[string]string{"a": "3", "b": "0"})
	handlers.Div(response, request)
	assert.Equal(t, "+Inf", response.Body.String())
}

func TestSqrt(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/sqrt/4", nil)
	request = mux.SetURLVars(request, map[string]string{"a": "4"})
	handlers.Sqrt(response, request)
	assert.Equal(t, "2", response.Body.String())
}

func TestCbrt(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/cbrt/27", nil)
	request = mux.SetURLVars(request, map[string]string{"a": "27"})
	handlers.Cbrt(response, request)
	assert.Equal(t, "3", response.Body.String())
}

func TestCbrtBadInput(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/cbrt/x", nil)
	request = mux.SetURLVars(request, map[string]string{"a": "x"})
	handlers.Cbrt(response, request)
	assert.Equal(t, "0", response.Body.String())
}

func TestCbrtBadRequest(t *testing.T) {
	response := httptest.NewRecorder()
	handlers.Cbrt(response, nil)
	assert.Equal(t, "0", response.Body.String())
}

func TestCbrtCorruptData(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "https://localhost/cbrt/0", nil)
	handlers.Cbrt(response, request)
	assert.Equal(t, "0", response.Body.String())
}

func TestGetRouterNotNil(t *testing.T) {
	var router = handlers.GetRouter()
	assert.NotNil(t, router)
}
