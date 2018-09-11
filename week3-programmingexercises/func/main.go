package main

import "fmt" 

type person struct {
	fName string
	lName string
}

func main() {
	p1 := person {
		fName: "Maria",
		lName: "Velasco",
	}

	fmt.Println(p1)

	fmt.Println("----------")

	f := foo()
	fmt.Println(f)
}


// func (receiver) identifier(parameters) (returns) { <code> }
func foo() int {
	return 34
}