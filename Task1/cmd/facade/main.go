package main

import (
	"fmt"
	"github.com/RyzhAlexWork/go-intern/Task1/pkg/models"

	"github.com/RyzhAlexWork/go-intern/Task1/pkg/facade"
	"github.com/RyzhAlexWork/go-intern/Task1/pkg/status"
	"github.com/RyzhAlexWork/go-intern/Task1/pkg/wallet"
)

var (
	expect = []models.Status{
		"Deposit was successful.",
		"Pay was successful.",
	}
	result  models.Status
	balance int
)

func main() {
	newWallet := wallet.NewWallet(0, 1000000)
	newStatus := walletstatus.NewWalletStatus()
	newUser := facade.NewUser("Aksel", newWallet, newStatus)
	result = newUser.Add(500000)
	if result != expect[0] {
		fmt.Printf("Expect result to equal: %s, but: %s.\n", expect, result)
	}
	result = newUser.Pay(200000)
	if result != expect[1] {
		fmt.Printf("Expect result to equal: %s, but: %s.\n", expect, result)
	}
	balance = newUser.Balance()
	if balance != 300000 {
		fmt.Printf("Wrong balance! Expect: 300000, balance: %d", balance)
	} else {
		fmt.Println("All ok!")
	}
}
