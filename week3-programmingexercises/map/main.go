package main

import "fmt"

func main() {
	m := map[string]int{"Maria": 34, "Jeanette": 27, "Zoraida": 22, "Jesenia": 17}
	fmt.Println(m)

	for k := range m {					// Or you can use -> for k, _ := range m
		fmt.Println(k)
	}

	for k, v := range m {
		fmt.Println(k, "-", v)
	}
}