package main

import (
	"net/http"
	"io"
	"fmt"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is from home")
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is from about")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is from contact")
}
