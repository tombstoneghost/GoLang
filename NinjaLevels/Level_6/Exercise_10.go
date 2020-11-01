package main

import "fmt"

func main() {
	f := fun()
	fmt.Println(f())
	fmt.Println(f())

	g := fun()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}

func fun() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
