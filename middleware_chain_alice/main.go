package main

import (
	"log"
	"net/http"

	"github.com/justinas/alice"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {
	finalHandler := http.HandlerFunc(final)
	chain := alice.New(middlewareOne, middlewareTwo)
	http.Handle("/", chain.Then(finalHandler))
	http.ListenAndServe(":3000", nil)
}
