package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
)

// START OMIT

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() { // HL
			if err := recover(); err != nil { // HL
				log.Println("Recover from error:", err)
				http.Error(w, http.StatusText(500), 500) // HL
			}
		}()
		log.Println("Executing recoveryMiddleware")
		next.ServeHTTP(w, r) // HL
	})

}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	panic("walau!")
	w.Write([]byte("OK"))
}

// END OMIT

func HitEndpoint(handler http.Handler, request *http.Request) (resp *httptest.ResponseRecorder) {
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, request)
	fmt.Println("Code: " + strconv.Itoa(resp.Code))
	fmt.Println("Body: " + resp.Body.String())
	return
}

func PanicInHandler() {
	finalHandler := http.HandlerFunc(final)
	request, _ := http.NewRequest("GET", "/", nil)
	fmt.Println("SIMULATION ===> Panic in Handler")
	HitEndpoint(finalHandler, request)
}

func RecoveryMiddleware() {
	finalHandler := http.HandlerFunc(final)
	request, _ := http.NewRequest("GET", "/", nil)
	fmt.Println("SIMULATION ===> With Recovery Middleware")
	HitEndpoint(recoveryMiddleware(finalHandler), request)
}

func main() {
	RecoveryMiddleware()
	// PanicInHandler()
}
