package main

import (
	"fmt"
)

type person struct {
	last string
	age  int
}

func main() {
	p1 := person{
		last: "Moneypenny", 
		age: 24,
	}
	fmt.Println(p1)
	p1.foo()

	p1.walk()
}

func (p person) foo() {
	fmt.Println("Hello from foo")
}

func (p person) walk() {
	fmt.Println(p.last, "is walking.")
}

// func (receiver) identifier(parameters) (returns) {}