package main 

import (
	"log"
	"os"
	"text/template"
)

type Book struct {
	Title string
	Author string
	Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	info := Book{
		Title: "Hunger Games",
		Author: "Suzanne Collins",
		Input: `<script>alert("Yow!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", info)
	if err != nil {
		log.Fatalln(err)
	}
}