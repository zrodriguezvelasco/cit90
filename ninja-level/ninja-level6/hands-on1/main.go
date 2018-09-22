package main 

import "fmt"

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x, y, z)
}

func foo() int {
	return 21
}

func bar() (int, string) {
	return 1996, "Blah blah blah"
}