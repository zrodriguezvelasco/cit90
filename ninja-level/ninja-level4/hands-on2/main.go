package main

import "fmt"

func main() {
	x := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

	for _, v := range x {
		fmt.Println(v)
		fmt.Printf("%T\n", v)
	}
}