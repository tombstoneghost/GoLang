package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	/*Method-2
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
	*/

	fmt.Println(<-c)
}
