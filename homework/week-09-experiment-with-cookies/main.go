package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/bowzer", bowzer)
	http.HandleFunc("/cat", cat)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var c *http.Cookie
	c, err := r.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}
	tpl.ExecuteTemplate(w, "index.html", c)
}

func bowzer(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:  "user-cookie",
		Value: "this would be the value",
		Path:  "/",
		//Path:  "/dog/bowzer/",
	}
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "bowzer.html", c)
}

func cat(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}
	tpl.ExecuteTemplate(w, "cat.html", c)
}