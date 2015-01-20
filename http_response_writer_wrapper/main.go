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

// START 3 OMIT

func specialLoggerMiddlewareWrapper(logger *AppLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			startTime := time.Now()

			w2 := &responseWriterLogger{w: w} // HL
			next.ServeHTTP(w2, r)             // HL

			logger.Info(r.Method + " " + r.URL.String() + " " + 
				time.Since(startTime).String() +
				" status: " + strconv.Itoa(w2.data.status) + // HL
				" size: " + strconv.Itoa(w2.data.size)) // HL

		})
	}
}

// END 3 OMIT

// START 1 OMIT

type responseWriterLogger struct {
	w    http.ResponseWriter
	data struct {
		status int
		size   int
	}
}

// END 1 OMIT

// START 2 OMIT

func (l *responseWriterLogger) Header() http.Header {
	return l.w.Header() // HL
}

func (l *responseWriterLogger) Write(b []byte) (int, error) {

	// scenario where WriteHeader has not been called
	if l.data.status == 0 {
		l.data.status = http.StatusOK
	}
	size, err := l.w.Write(b) // HL
	l.data.size += size
	return size, err // HL
}

func (l *responseWriterLogger) WriteHeader(code int) {
	l.w.WriteHeader(code) // HL
	l.data.status = code
}

// END 2 OMIT

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

	specialLoggerMiddleware := specialLoggerMiddlewareWrapper(NewLogger())

	// finalHandler := http.HandlerFunc(final)
	// http.Handle("/", specialLoggerMiddleware(finalHandler))
	// http.ListenAndServe(":3000", nil)

	finalHandler := http.HandlerFunc(final)
	request, _ := http.NewRequest("GET", "/", nil)
	HitEndpoint(specialLoggerMiddleware(finalHandler), request) // HL

}
