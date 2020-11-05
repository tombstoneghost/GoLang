package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send
	go send(eve, odd, quit)

	// Receive
	receive(eve, odd, quit)

	fmt.Println("Exiting...")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From even Channel:", v)
		case v := <-o:
			fmt.Println("From odd Channel:", v)
		case v := <-q:
			fmt.Println("From quit Channel:", v)
			return

		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
