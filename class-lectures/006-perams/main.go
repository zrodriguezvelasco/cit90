package main

import (
	"fmt"
)

func main() {
	fmt.Println("begin")

	// pass in an ARGUMENT to the func when you CALL IT
	foo("James Bond")

	
	x := "moneypenny"
	// pass in an ARGUMENT to the func when you CALL IT
	foo(x)

	fmt.Println("about to end")
}

// foo takes a VALUE of TYPE string
// we could also say;
// foo has a parameter which is a VALUE of TYPE string
// funcs may be defined with PARAMETER(s)
func foo(name string) {
	fmt.Println("hello", name)
}