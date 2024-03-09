package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /variables", variables)
	mux.HandleFunc("POST /snippets", createSnippets)
	fmt.Println("Listing server on http://localhost:4000")
	http.ListenAndServe(":4000", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hellow from root"))
}

func variables(w http.ResponseWriter, r *http.Request) {
	// type conversion
	byteSlice := []byte{'H', 'e', 'l', 'l', 'o'}
	byteSlice1 := []byte("Hellow from byte1")
	fmt.Println("Slice of bytes:", byteSlice, byteSlice1)
	fmt.Println("byte to string:", string(byteSlice), string(byteSlice1))
	w.Write([]byte("check the output in terminal"))
}

func createSnippets(w http.ResponseWriter, r *http.Request) {
	// this if check is useless now as we have added method POST in the routing as part of go 1.22
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// If the method is POST, execute this block
	w.Write([]byte("Create a new snippet..."))
}
