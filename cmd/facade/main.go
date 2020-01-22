package main

import (
	"fmt"

	"github.com/RyzhAlexWork/go-intern/pkg/pay"
	"github.com/RyzhAlexWork/go-intern/pkg/status"
	"github.com/RyzhAlexWork/go-intern/pkg/facade"
	"github.com/RyzhAlexWork/go-intern/pkg/wallet"
)

var (
	expect = "Wallet was create.\nStatus: Nice day! :)\nPay was success."
)

func main() {
	wallet := wallet.NewWallet()
	status := status.NewStatus()
	pay := pay.NewPay(wallet,1000)
	user := facade.NewUser(wallet, status, pay)
	result := user.FirstSignIn()

	if result != expect {
		fmt.Println("Expect result to equal %s, but %s.\n", expect, result)
	} else {
		fmt.Println("All ok! :)")
	}
}
