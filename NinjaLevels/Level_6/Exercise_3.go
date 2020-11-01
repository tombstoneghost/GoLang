package main

import "fmt"

func main() {
	defer foo3()
	bar3()
}

func foo3() {
	fmt.Println("foo")
}

func bar3() {
	fmt.Println("bar")
}
