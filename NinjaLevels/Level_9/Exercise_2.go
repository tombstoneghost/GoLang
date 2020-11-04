package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello, from", p.first)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{first: "James"}

	saySomething(&p1)

	p1.speak()
}
