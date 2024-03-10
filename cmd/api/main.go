package main

import (
	"flag"
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
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	fmt.Println("Listing server on http://localhost", *addr)
	http.ListenAndServe(*addr, mux)
}
