package main

import (
	"fmt"

	"github.com/RyzhAlexWork/go-intern/Facade/pkg/facade"
)

func main() {
	expect := "Wallet was create.\nStatus: Nice day! :)\nPay was success."

	user := facade.NewUser()
	result := user.FirstSignIn()

	if result != expect {
		fmt.Println("Expect result to equal %s, but %s.\n", expect, result)
	} else {
		fmt.Println("All ok! :)")
	}
}
