package main

import "fmt"

type person3 struct {
	first string
	last  string
}

type secretAgent2 struct {
	person3
	ltk bool
}

func (s secretAgent2) speak() {
	fmt.Println("I'm", s.first, s.last, "A Secret Agent.")
}

func (p person3) speak() {
	fmt.Println("I'm", p.first, p.last, "A Normal Person")
}

type human interface {
	speak()
}

func bar3(h human) {
	switch h.(type) {
	case person3:
		fmt.Println("I was passed into barr", h.(person3).first, "A Person.")
	case secretAgent2:
		fmt.Println("I was passed into bar", h.(secretAgent2).first, "An Agent")
	}
}

type hotdog int

func main() {
	sa1 := secretAgent2{
		person3: person3{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent2{
		person3: person3{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}

	p1 := person3{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar3(sa1)
	bar3(sa2)
	bar3(p1)

	// Conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
