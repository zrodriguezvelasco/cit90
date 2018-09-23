package main 

import (
	"os"
	"text/template"
	"log"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal("whoops", err)
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "James")
}