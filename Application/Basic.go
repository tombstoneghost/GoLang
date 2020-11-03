package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

// For JSON to struct conversion
type personJSON struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	// Converting to JSON
	fmt.Println("Struct Slice to JSON:")
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	// Converting JSON to Go Data Structure
	fmt.Println("\nJSON to Struct Slice")
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs = []byte(s)

	var peopleJSON []personJSON

	err = json.Unmarshal(bs, &peopleJSON)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(peopleJSON)
}
