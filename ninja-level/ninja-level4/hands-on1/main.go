package main

import "fmt"

func main() {
	x := [5]int{5, 10, 15, 20, 25}

	for _, v := range x {
		fmt.Println(v)
		fmt.Printf("%T\n", v)
	}
}