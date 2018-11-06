package main 

import "fmt"

func main() {
	fmt.Println("Hello World")
	secondStatement()
	finalState()
}

// func receiver identifier(parameters) returns {code}

func secondStatement() {
	fmt.Println("Here is my second statement")
}

func finalState() {
	fmt.Println("about to exit")
}