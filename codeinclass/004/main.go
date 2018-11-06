package main

import "fmt"

func main() {
	z := multiply(5, 5)
	fmt.Println(z)
}

// func receiver indentifier(parameter) return {code}
func multiply(x int, y int) int {
	return x * y
}



