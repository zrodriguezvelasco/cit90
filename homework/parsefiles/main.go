package main

import (
	"text/template"
	"log"
	"os"
	"fmt"
)

func main() {
	tpl, err := template.ParseFiles("one.txt", "two.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}
	tpl.ExecuteTemplate(os.Stdout, "one.txt", "blah blah blah")

	fmt.Println("\n-----")
	tpl.ExecuteTemplate(os.Stdout, "two.txt", "Zoraida")
}