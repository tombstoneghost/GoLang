package main

import "fmt"

type newType int

var k newType

func main()  {
	fmt.Println(k)
	fmt.Printf("%T\n", k)
	k = 42
	fmt.Println(k)
}