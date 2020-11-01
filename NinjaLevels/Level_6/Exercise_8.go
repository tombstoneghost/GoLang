package main

import "fmt"

func main() {
	f := number()

	fmt.Println(f())
}

func number() func() int {
	return func() int {
		return 42
	}
}
