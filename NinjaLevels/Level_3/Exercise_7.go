package main

import "fmt"

func main() {
	x := 10

	if x == 10 {
		fmt.Println(x)
	} else if x == 11 {
		fmt.Println("x is 11, not 10")
	} else {
		fmt.Printf("x is unknown")
	}
}
