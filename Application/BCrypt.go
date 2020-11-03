package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password1234"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(password)
	fmt.Println(bs)

	loginPassword := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))

	if err != nil {
		fmt.Println("Can't Login")
		return
	}

	fmt.Println("Logged In")
}
