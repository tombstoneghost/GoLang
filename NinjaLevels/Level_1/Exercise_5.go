package main

import "fmt"

type newType2 int

var p newType2
var q int

func main()  {
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	p = 42
	fmt.Println(p)

	q = int(p)
	fmt.Println(q)
	fmt.Printf("%T\n", q)
}