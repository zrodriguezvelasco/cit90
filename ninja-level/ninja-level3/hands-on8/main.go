package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Not printing")
	case true:
		fmt.Println("Printing")
	}
}