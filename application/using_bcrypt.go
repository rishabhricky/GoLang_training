package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	x := "helloworld@1234"
	bs, err := bcrypt.GenerateFromPassword([]byte(x), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(x)
	fmt.Println(bs)

	loginPwd := "helloworld@12345678"


	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd))

	if err != nil {
		fmt.Println("You can't login")
		return
	}

	fmt.Println("You're logged in")
}
