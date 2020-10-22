package main

import "fmt"

func main()  {
	a := (42 == 42)
	fmt.Println(a)

	a = (42 <= 43)
	fmt.Println(a)

	a = (42 >= 43)
	fmt.Println(a)

	a = (42 != 43)
	fmt.Println(a)

	a = (42 < 43)
	fmt.Println(a)

	a = (42 > 43)
	fmt.Println(a)
}