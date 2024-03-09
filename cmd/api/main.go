package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/variables", variables)
	http.ListenAndServe(":4000", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
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
