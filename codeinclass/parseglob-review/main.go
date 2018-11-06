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
	//Execute the main.gohtml file
	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", nil)
	if err != nil {
		log.Fatal("oops", err)
	}

	//Execute the about.gohtml file
	err = tpl.ExecuteTemplate(os.Stdout, "about.gohtml", nil)
	if err != nil {
		log.Fatal("oops", err)
	}

	//Create the f file called main.html
	f, err := os.Create("main.html")
	if err != nil {
		log.Fatal("oops", err)
	}
	defer f.Close()

	//Create the f2 file called about.html
	f2, err := os.Create("about.html")
	if err != nil {
		log.Fatal("oops", err)
	}
	defer f2.Close()

	// Execute the f file and copy the main.gohtml onto the f file
	err = tpl.ExecuteTemplate(f, "main.gohtml", nil)
	if err != nil {
		log.Fatal("oops", err)
	}

	// Execute the f2 file and copy the main.gohtml onto the f2 file
	err = tpl.ExecuteTemplate(f2, "about.gohtml", nil)
	if err != nil {
		log.Fatal("oops", err)
	}
}