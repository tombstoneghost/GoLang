package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("Square: ", x*x)
	}(5)
}
