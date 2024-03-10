package main

import (
	"fmt"
	"net/http"

	handler "github.com/JayeshKarandeDev/microservice/cmd/api/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.Home)
	mux.HandleFunc("GET /variables", handler.Variables)
	mux.HandleFunc("POST /snippets", handler.CreateSnippets)
	mux.HandleFunc("GET /snippet/view", handler.ViewSnippet)
	fmt.Println("Listing server on http://localhost:4000")
	http.ListenAndServe(":4000", mux)
}
