package main

import "fmt"

func main() {
	sq := func(x int, n int) int {
		res := 1
		for i := 0; i < n; i++ {
			res = res * x
		}

		return res
	}

	fmt.Println("2 raised to power 8 is:", sq(2, 8))

}
