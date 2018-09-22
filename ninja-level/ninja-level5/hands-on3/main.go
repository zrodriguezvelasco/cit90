package main 

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle 
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println(t.color)
	fmt.Println(s.color)
}