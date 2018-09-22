package main

import "fmt"

func main() {
	a := (5 == 5)
	b := (5 <= 6)
	c := (5 >= 6)
	d := (5 != 6)
	e := (5 < 6)
	f := (5 > 6)
	fmt.Println(a, b, c, d, e, f)
}