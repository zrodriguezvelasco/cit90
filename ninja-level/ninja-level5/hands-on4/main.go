package main 

import "fmt"

func main() {
	s := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "Zoraida",
		friends: map[string]int {
			"Maria": 34,
			"Jeanette": 27,
		},
		favDrinks: []string {
			"Tea",
			"Coke",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	fmt.Println("----------")

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	fmt.Println("----------")

	for i, v := range s.favDrinks {
		fmt.Println(i, v)
	}
}