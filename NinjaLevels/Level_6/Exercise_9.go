package main

import "fmt"

func main() {
	f := func(xi []int) int {
		sum := 0

		for _, v := range xi {
			sum += v
		}

		return sum
	}

	x := sum(f, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	fmt.Println(x)
}

func sum(f func(xi []int) int, ii []int) int {
	n := f(ii)

	return n
}
