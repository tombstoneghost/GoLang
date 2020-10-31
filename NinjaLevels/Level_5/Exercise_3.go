package main

import "fmt"

type vehicle struct {
	doors string
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
	truck1 := truck{
		vehicle: vehicle{
			doors: "two",
			color: "red",
		},
		fourWheel: true,
	}

	car := sedan{
		vehicle: vehicle{
			doors: "four",
			color: "gray",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(car)

	fmt.Println(truck1.doors)
	fmt.Println(car.doors)
}
