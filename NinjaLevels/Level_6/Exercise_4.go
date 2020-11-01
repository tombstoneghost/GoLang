package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello", p.first, p.last, "You are", p.age, "old.")
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p.speak()
}
