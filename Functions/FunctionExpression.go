package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("First Function Expression")
	}

	f()

	g := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}

	g(1942)
}
