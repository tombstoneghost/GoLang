package main

import "fmt"

func main() {
	// Anonymous Function

	func() {
		fmt.Println("Anonymous Function Ran.")
	}()

	func(x int) {
		fmt.Println("Number Passed:", x)
	}(42)
}
