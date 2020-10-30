package main

import "fmt"

func main() {
	mapData := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice Cream", "Sunsets"},
	}

	mapData["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for k, v := range mapData {
		fmt.Println("Record: ", k)

		for i, val := range v {
			fmt.Printf("\t%v\t%v\n", i, val)
		}
	}
}
