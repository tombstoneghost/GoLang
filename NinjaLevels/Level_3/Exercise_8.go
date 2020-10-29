package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This is a false statement")
	case true:
		fmt.Printf("This is a true statement\n")
	}
}
