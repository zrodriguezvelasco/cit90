package main 

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", ind)
	http.HandleFunc("/about", abo)

	http.ListenAndServe(":8080", nil)
}

func ind(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", 42)
}

func abo(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}