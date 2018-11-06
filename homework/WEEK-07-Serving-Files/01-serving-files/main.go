package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	http.HandleFunc("/kitten.jpg", kitten)
	http.HandleFunc("/hamster.jpg", hamsterPic)
	http.HandleFunc("/panda.jpg", pandaPic)
	http.HandleFunc("/penguin.jpg", penguinPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

func kitten(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "kitten.jpg")
}

func hamsterPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "hamster.jpg")
}

func pandaPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "panda.jpeg")
}

func penguinPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "penguin.jpg")
}