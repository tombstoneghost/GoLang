package main

import "fmt"

var y = 43

// This means that the variable named as 'z' is of type 'int'
var z int

func main()  {
	// Using var keyword to declare variable
	var x = 42
	fmt.Println(x)

	fmt.Println(y)

	// Default value of z is 0
	fmt.Println(z)

	/* := operator cannot be used outside the function scope*/
}
