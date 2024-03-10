package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Hellow from root"))
	files := []string{
		"ui/html/pages/home.tmpl.html",
		"ui/html/pages/base.tmpl.html",
		"ui/html/pages/nav.tmpl.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}
	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// log.print(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}

}

func Variables(w http.ResponseWriter, r *http.Request) {
	// type conversion
	byteSlice := []byte{'H', 'e', 'l', 'l', 'o'}
	byteSlice1 := []byte("Hellow from byte1")
	fmt.Println("Slice of bytes:", byteSlice, byteSlice1)
	fmt.Println("byte to string:", string(byteSlice), string(byteSlice1))
	w.Write([]byte("check the output in terminal"))
}

func CreateSnippets(w http.ResponseWriter, r *http.Request) {
	// this if check is useless now as we have added method POST in the routing as part of go 1.22
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// If the method is POST, execute this block
	w.Write([]byte("Create a new snippet..."))
}

func ViewSnippet(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("id")
	fmt.Println("id is", id1)
	fmt.Printf("type is %T", id1)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
