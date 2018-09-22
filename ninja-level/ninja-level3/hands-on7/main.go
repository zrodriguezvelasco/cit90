package main 

import "fmt"

func main() {
	x := 5

	if x == 6 {
		fmt.Println(x)
	} else if x == 5 {
		fmt.Println(true)
	} else {
		fmt.Println("none")
	}
}