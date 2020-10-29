package main

import "fmt"

func main() {
	/* if else if else */
	// Basics
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}

	// Initialization
	if x := 42; x == 2 {
		fmt.Println("001")
	}

	// if-else
	x := 42
	if x == 40 {
		fmt.Println("Our value was 40")
	} else {
		fmt.Println("Our value was 42")
	}

	// if-else if-else
	x = 41
	if x == 40 {
		fmt.Println("Our value was 40")
	} else if x == 41 {
		fmt.Println("Our value was 41")
	} else {
		fmt.Println("Our value was 42")
	}

	/* Switch */
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond":
		fmt.Println("Miss Money or Bond")
	case "M":
		fmt.Println("M")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("This is default")
	}

	// Boolean Operators
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
