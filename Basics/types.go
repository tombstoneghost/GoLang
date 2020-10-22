package main

import "fmt"

var i = 42
var s = "Hello World!"

func main()  {
	fmt.Println(i)
	fmt.Printf("%T\n", i) // Prints the type of i

	fmt.Println(s)
	fmt.Printf("%T\n", s) // Prints the type of s
}