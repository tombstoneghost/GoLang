package main

import "fmt"

func main() {
	num := 42
	fmt.Printf("%d\t%b\t%#x\n", num, num, num)

	num2 := num << 1
	fmt.Printf("%d\t%b\t%#x\n", num2, num2, num2)
}
