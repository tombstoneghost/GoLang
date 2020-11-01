package main

import "fmt"

func main() {
	n := factorial(4)

	fmt.Println("Factorial of 4:", n)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
