package main

import "net/http"

func handlerFn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello world!`))
}

func main() {
	http.HandleFunc("/", handlerFn)
	http.ListenAndServe("localhost:4000", nil)
}
