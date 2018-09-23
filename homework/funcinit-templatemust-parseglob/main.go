package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		panic(err)
	}
}