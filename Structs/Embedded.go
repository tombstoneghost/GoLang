package main

import "fmt"

type person2 struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person2
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person2: person2{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person2{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
}
