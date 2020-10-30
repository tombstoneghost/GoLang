package main

import "fmt"

func main() {
	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss", "Moneypenny", "Hello, James"}

	dict := [][]string{slice1, slice2}

	for i, v := range dict {
		fmt.Println("Record: ", i)

		for j, val := range v {
			fmt.Printf("\tIndex: %v\tValue:%v\n", j, val)
		}
	}
}
