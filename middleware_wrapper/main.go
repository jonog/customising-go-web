package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"time"
)

type AppLogger struct {
	logger *log.Logger
}

func NewLogger() *AppLogger {
	return &AppLogger{logger: log.New(os.Stdout, "==> ", log.Ldate|log.Ltime)}
}

func (a *AppLogger) Info(msg string) {
	a.logger.Println("INFO: ", msg)
}

// START OMIT

func simpleLoggerMiddlewareWrapper(logger *AppLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler { // HL
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			startTime := time.Now()

			next.ServeHTTP(w, r)

			endTime := time.Since(startTime)
			logger.Info(r.Method + " " + r.URL.String() + " " + endTime.String())
		})
	}
}

// END OMIT

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
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

	simpleLoggerMiddleware := simpleLoggerMiddlewareWrapper(NewLogger())

	finalHandler := http.HandlerFunc(final)

	// http.Handle("/", simpleLoggerMiddleware(finalHandler))
	// http.ListenAndServe(":3000", nil)

	request, _ := http.NewRequest("GET", "/", nil)
	HitEndpoint(simpleLoggerMiddleware(finalHandler), request) // HL
}
