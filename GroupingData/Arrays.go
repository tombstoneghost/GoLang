package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[3] = 42
	fmt.Println(arr)
	fmt.Printf("Lenght of Array: %d", len(arr))

}
