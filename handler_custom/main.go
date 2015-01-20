package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"
)

// START 1 OMIT

type appHandler struct {
	h func(http.ResponseWriter, *http.Request) (error) // HL
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := ah.h(w, r)
	if err != nil {
		switch err := err.(type) { // HL
		case ErrorDetails: // HL
			ErrorJSON(w, err) // HL
		default:
			ErrorJSON(w, ErrorDetails{"Internal Server Error", "", 500})
		}
	}
}

// END 1 OMIT

// START 2 OMIT
type ErrorDetails struct {
	Message string `json:"error"`
	Details string `json:"details,omitempty"`
	Status int `json:"-"`
}

func (e ErrorDetails) Error() string {
	return fmt.Sprintf("Error: %s, Details: %s", e.Message, e.Details)
}

func ErrorJSON(w http.ResponseWriter, details ErrorDetails) {

	jsonB, err := json.Marshal(details)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json") // HL
	w.WriteHeader(details.Status) // HL
	w.Write(jsonB) // HL
}

// END 2 OMIT

// START 3 OMIT
func unstableEndpoint(w http.ResponseWriter, r *http.Request) (error) {

	if rand.Intn(100) > 60 {
		return ErrorDetails{"Strange request", "Please try again.", 422} // HL
	}

	if rand.Intn(100) > 80 {
		return ErrorDetails{"Serious failure", "We are investigating.", 500} // HL
	}

	w.Write([]byte(`{"ok":true}`))
	return nil
}

// END 3 OMIT

// SECTION2

func HitEndpoint(handler http.Handler, request *http.Request) (resp *httptest.ResponseRecorder) {
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, request)
	fmt.Println("Code: " + strconv.Itoa(resp.Code))
	fmt.Println("Body: " + resp.Body.String())
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// http.Handle("/", appHandler{unstableEndpoint})
	// http.ListenAndServe(":3000", nil)

	request, _ := http.NewRequest("GET", "/", nil)
	for i := 0; i < 8; i++ {
		time.Sleep(500*time.Millisecond)
		HitEndpoint(appHandler{unstableEndpoint}, request)
	}

}
