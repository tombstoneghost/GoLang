package main

import "fmt"

func main() {
	c := make(chan int, 1) // Change Value from 1 to 2 inorder to correct the unsuccessful channel

	// Won't Work
	// c <- 42

	go func() {
		// Successful Buffer
		c <- 42

		// Unsuccessful Buffer
		c <- 43
	}()

	fmt.Println(<-c)
}
