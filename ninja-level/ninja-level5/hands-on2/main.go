package main 

import "fmt"

type person struct {
	fName string
	lName string
	favIceCreamFlavors []string
}

func main() {
	p1 := person {
		fName: "Maria",
		lName: "Velasco",
		favIceCreamFlavors: []string{"Chocolate chip", "Mint"},
	}

	p2 := person {
		fName: "Jeanette",
		lName: "Castro",
		favIceCreamFlavors: []string{"Vanilla", "Chocolate chip mint"},
	}

	fmt.Println(p1.fName)
	fmt.Println(p1.lName)
	for i, v := range p1.favIceCreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.fName)
	fmt.Println(p2.lName)
	for i, v := range p2.favIceCreamFlavors {
		fmt.Println(i, v)
	}

	m := map[string]person {
		p1.lName: p1,
		p2.lName: p2,
	}

	for _, v := range m {
		fmt.Println(v)
	}
}