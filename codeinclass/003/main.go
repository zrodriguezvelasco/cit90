package main

import "fmt"

func main() {
	// when we call a func, we pass in ARGUMENTS
	// Calling foo function
	// Multiple arguments
	foo("Zoraida", 21)
	foo("Frank", 40)
}
// func
// modularize our code and DRY (don't repeat yourself)
// func are defined w/ parameters
// parameters specify a VALUE of a certain TYPE to be passed in when the func is called
// 			Parameter
func foo(x string, y int) {
	fmt.Println(x, y)
}