package main

import (
	"fmt"
	"sort"
)

type userData struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []userData

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByLast []userData

func (l ByLast) Len() int {
	return len(l)
}

func (l ByLast) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ByLast) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

func main() {
	u1 := userData{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := userData{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := userData{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []userData{u1, u2, u3}

	fmt.Println(users)

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)

		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)

		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	sort.Sort(ByLast(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)

		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

}
