package main

import "fmt"

type person2 struct {
	first     string
	last      string
	favorites []string
}

func main() {
	p1 := person2{
		first:     "James",
		last:      "Bond",
		favorites: []string{"Butterscotch", "Strawberry"},
	}

	p2 := person2{
		first:     "Miss",
		last:      "Moneypenny",
		favorites: []string{"Chocolate", "Vanilla"},
	}

	m := map[string]person2{
		p1.last: p1,
		p2.last: p2,
	}

	for i, v := range m {
		fmt.Printf("%v %v\n", v.first, i)

		for j, val := range v.favorites {
			fmt.Println(j+1, val)
		}

		fmt.Println("----------------------")
	}
}
