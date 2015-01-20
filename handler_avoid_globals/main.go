package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
)

// START OMIT

type Api struct {
	importantThing string // HL
	// db *gorp.DbMap
	// redis *redis.Pool
	// logger ...
}

type appHandler struct {
	*Api
	h func(*Api, http.ResponseWriter, *http.Request)
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.h(ah.Api, w, r)
}

func myHandler(a *Api, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("2015: Year of the " + a.importantThing)) // HL
}

// END OMIT

func HitEndpoint(handler http.Handler, request *http.Request) (resp *httptest.ResponseRecorder) {
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, request)
	fmt.Println("Code: " + strconv.Itoa(resp.Code))
	fmt.Println("Body: " + resp.Body.String())
	return
}

func main() {

	context := &Api{importantThing: "Gopher"}
	// http.Handle("/", appHandler{context, myHandler})
	// http.ListenAndServe(":3000", nil)

	request, _ := http.NewRequest("GET", "/", nil)
	HitEndpoint(appHandler{context, myHandler}, request)
}
