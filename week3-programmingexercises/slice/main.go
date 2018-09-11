package main

import "fmt"

func main() {
	s := []int{7, 8, 9, 46}
	fmt.Println(s)

	for idx := range s {					// Or can use -> for idx, _ := range s
		fmt.Println(idx)
	}

	for idx, v := range s {
		fmt.Println(idx, "-", v)
	}
}