package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Basic Error Checking

	fmt.Println("Basic Error Checking:")
	n, err := fmt.Println("Hello")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}

	// Checking Input Errors

	fmt.Println("Checking Input Errors:")

	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav. Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav. Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)

	// File Creating Error

	fmt.Println("File Creating Errors:")

	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")

	_, _ = io.Copy(f, r)

	// Opening and Reading a File Errors

	fmt.Println("Opening and Reading a File Errors")

	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
