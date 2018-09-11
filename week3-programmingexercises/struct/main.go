package main

import "fmt" 

type person struct {
	fName string
	lName string
	favFood []string
}

func main() {
	p1 := person {
		fName: "Maria",
		lName: "Velasco",
		favFood: []string {"Mexican rice", "Birria", "Eggs"},
	}

	fmt.Println(p1)

	fmt.Println("----------")

	fmt.Println(p1.fName)

	fmt.Println("----------")

	fmt.Println(p1.favFood)

	fmt.Println("----------")

	for _, a := range p1.favFood {
		fmt.Println(a)
	} 
}