package main

import "fmt"

func main() {
	xs := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum := variadic(xs...)
	fmt.Println("Total is:", sum)

	// Passing 0 Argument
	sum = variadic()
	fmt.Println("Zero Argument Result:", sum)
}

func variadic(x ...int) int {
	fmt.Println(x)

	sum := 0

	for i, v := range x {
		fmt.Println("For item in index position", i, "we are adding", v, "to the total")
		sum += v
	}

	return sum
}
