package main

import (
	"log"
	"text/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	// Execute the f file and copy the main.gohtml onto the f file
	f1, err := tpl.ExecuteTemplate(f1, "main.gohtml", 42)
	if err != nil {
		log.Fatal("oops", err)
	}
}
