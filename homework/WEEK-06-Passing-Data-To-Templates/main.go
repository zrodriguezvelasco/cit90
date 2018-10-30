package main 

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/about", bar)
	http.HandleFunc("/contact", con)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", 42)
}

func bar(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", []int{1,2,3,4,5})
}

func con(w http.ResponseWriter, r *http.Request) {
	type person struct {
		First string
		Age int
	}

	p1 := person {
		First: "Zoraida",
		Age: 22,
	}

	p2 := person {
		First: "Jeanette",
		Age: 27,
	}

	xp := []person{p1, p2}

	tpl.ExecuteTemplate(w, "contact.html", xp)
}