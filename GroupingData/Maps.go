package main

import (
	"fmt"
)

func main() {
	// Basic
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])

	// Adding to Map
	m["Todd"] = 33
	fmt.Println(m)

	// Range
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Deleting
	delete(m, "Todd")
	fmt.Println(m)
}
