package main

import "fmt"

type person struct {
	first     string
	last      string
	favorites []string
}

func main() {
	p1 := person{
		first:     "James",
		last:      "Bond",
		favorites: []string{"Butterscotch", "Strawberry"},
	}

	p2 := person{
		first:     "Miss",
		last:      "Moneypenny",
		favorites: []string{"Chocolate", "Vanilla"},
	}

	fmt.Println(p1.first, p1.last)

	for i, v := range p1.favorites {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.first, p2.last)

	for i, v := range p2.favorites {
		fmt.Println("\t", i, v)
	}
}
