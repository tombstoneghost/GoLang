package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	res1 := foo2(x...)
	fmt.Println(res1)

	res2 := bar2(x)
	fmt.Println(res2)
}

func foo2(n ...int) int {
	sum := 0

	for _, v := range n {
		sum += v
	}

	return sum
}

func bar2(n []int) int {
	sum := 0

	for _, v := range n {
		sum += v
	}

	return sum
}
