package main

import "fmt"

func main() {
	// Creating a Slice
	x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(x)

	// Basic Slice Functions
	fmt.Println(len(x))
	fmt.Println(x[2])

	// Looping
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Slicing a Slice
	fmt.Println(x[1:3])

	// Appending to Slice
	x = append(x, 100, 45, 36)
	fmt.Println(x)

	y := []int{1, 2, 3, 4, 5}
	x = append(x, y...)
	fmt.Println(x)

	// Deleting from Slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// Creating slice using make
	intSlice := make([]int, 10, 100)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

	// Multi-dimensional Slice
	newSlice := [][]int{x, y}
	fmt.Println(newSlice)
}
