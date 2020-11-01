package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)
	x++

	{
		y := 42
		fmt.Println(y)

		fmt.Println(x)
	}

	// Undefined Scope
	//fmt.Println(y)

	fmt.Println(x)
	foo4()

	a := incrementer()
	b := incrementer()

	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
}

func foo4() {
	fmt.Println("Hello")
}

func incrementer() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
