package main 

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5}
	ix := foo(ii...)
	fmt.Println(ix)

	ii2 := []int{1, 2, 3, 4, 5}
	ix2 := bar(ii2)
	fmt.Println(ix2)
}

func foo(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func bar (i []int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}