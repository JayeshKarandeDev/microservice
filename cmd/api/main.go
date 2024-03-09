package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	http.ListenAndServe(":4000", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow from root"))
}
