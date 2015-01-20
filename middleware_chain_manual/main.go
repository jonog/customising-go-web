package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
)

// START 1 OMIT

func middlewareOne(next http.Handler) http.Handler { // HL
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("-> Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("-> Executing middlewareOne again")
	})
}

// END 1 OMIT
// START 2 OMIT

func middlewareTwo(next http.Handler) http.Handler { // HL
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("---> Executing middlewareTwo")
		next.ServeHTTP(w, r)
		log.Println("---> Executing middlewareTwo again")
	})
}

// END 2 OMIT

func final(w http.ResponseWriter, r *http.Request) { // HL
	log.Println("-----> Executing finalHandler")
	w.Write([]byte("OK"))
}

func HitEndpoint(handler http.Handler, request *http.Request) (resp *httptest.ResponseRecorder) {
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, request)
	fmt.Println("Code: " + strconv.Itoa(resp.Code))
	fmt.Println("Body: " + resp.Body.String())
	return
}

func main() {
	finalHandler := http.HandlerFunc(final)
	// http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
	// http.ListenAndServe(":3000", nil)

	request, _ := http.NewRequest("GET", "/", nil)
	fmt.Println("SIMULATION ===> Panic in Handler")
	HitEndpoint(middlewareOne(middlewareTwo(finalHandler)), request)

}
