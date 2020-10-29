package main

import "fmt"

func main() {
	favSport := "cricket"

	switch favSport {
	case "badminton":
		fmt.Println("Go to Badminton Court")
	case "basketball":
		fmt.Println("Go to Basketball Court")
	case "cricket":
		fmt.Println("Go to Ground")
	default:
		fmt.Println("Sit down and study Go")
	}
}
