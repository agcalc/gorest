package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetRouter returns the router configuration
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", Hello).Methods("GET")
	router.HandleFunc("/add/{a}/{b}", Add).Methods("GET")
	router.HandleFunc("/sub/{a}/{b}", Sub).Methods("GET")
	router.HandleFunc("/mult/{a}/{b}", Mult).Methods("GET")
	router.HandleFunc("/div/{a}/{b}", Div).Methods("GET")
	router.HandleFunc("/sqrt/{a}", Sqrt).Methods("GET")
	router.HandleFunc("/cbrt/{a}", Cbrt).Methods("GET")
	return router
}

// getFloat64ArgFromRequest will convert a request arg to a float or set it to 0 if there was some problem with the input
func getFloat64ArgFromRequest(arg string, request *http.Request) float64 {
	var (
		afloat float64
		err    error
	)

	if request == nil {
		return 0
	}
	if astring, ok := mux.Vars(request)[arg]; ok {
		afloat, err = strconv.ParseFloat(astring, 64)
	} else {
		afloat = 0
	}
	if err != nil {
		afloat = 0
	}

	return afloat
}

func generateSuccessFloatResponse(val float64, response http.ResponseWriter) {
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(strconv.FormatFloat(val, 'f', -1, 64)))
}

// Hello is the response for /
func Hello(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Hello World!"))
}

// Add adds two numbers
func Add(response http.ResponseWriter, request *http.Request) {
	a := getFloat64ArgFromRequest("a", request)
	b := getFloat64ArgFromRequest("b", request)

	generateSuccessFloatResponse(a+b, response)
}

// Sub is the difference between two numbers
func Sub(response http.ResponseWriter, request *http.Request) {
	a := getFloat64ArgFromRequest("a", request)
	b := getFloat64ArgFromRequest("b", request)

	generateSuccessFloatResponse(a-b, response)
}

// Mult is the multiplication of two numbers
func Mult(response http.ResponseWriter, request *http.Request) {
	a := getFloat64ArgFromRequest("a", request)
	b := getFloat64ArgFromRequest("b", request)

	generateSuccessFloatResponse(a*b, response)
}

// Div is the division of two numbers
func Div(response http.ResponseWriter, request *http.Request) {
	a := getFloat64ArgFromRequest("a", request)
	b := getFloat64ArgFromRequest("b", request)

	generateSuccessFloatResponse(a/b, response)
}

// Sqrt is the square root of a number
func Sqrt(response http.ResponseWriter, request *http.Request) {
	a := getFloat64ArgFromRequest("a", request)

	generateSuccessFloatResponse(math.Sqrt(a), response)
}

// Cbrt is the cubed root of a number
func Cbrt(response http.ResponseWriter, request *http.Request) {
	a := getFloat64ArgFromRequest("a", request)

	generateSuccessFloatResponse(math.Cbrt(a), response)
}
