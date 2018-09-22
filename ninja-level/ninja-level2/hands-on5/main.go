package main

import "fmt"

func main() {
	x := `blah blah
	some 
	raw literal string
	"fun fun"`
	fmt.Println(x)
}