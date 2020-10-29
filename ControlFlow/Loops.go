package main

import "fmt"

func main() {
	// Normal Loop
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	// Nested Loops
	for i := 0; i <= 10; i++ {
		fmt.Printf("\nOuter Loop: %d", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\n\t\tInner Loop: %d", j)
		}
	}

	// While loop alternative
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Internal Looping
	x = 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}

	// Use of break and continue
	x = 0
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
