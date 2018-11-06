package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is my doggie")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is my cat")
}

func main() {
	// Handle the root(name it), then input handlerfunc()
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))
	// Get localhost, put number
	http.ListenAndServe(":8080",nil)
}
