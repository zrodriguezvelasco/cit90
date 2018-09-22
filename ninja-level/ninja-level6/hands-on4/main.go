package main 

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println("i am", p.first, p.last, "and", p.age)
}

func main() {
	p1 := person{
		first: "Zoraida",
		last: "Rodriguez",
		age: 22,
	}

	p1.speak()
}