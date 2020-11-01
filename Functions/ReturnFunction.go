package main

import "fmt"

func main() {
	s1 := foo3()
	fmt.Println(s1)

	x := bar4()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

}

func foo3() string {
	s := "Hello World"
	return s
}

func bar4() func() int {
	return func() int {
		return 451
	}
}
