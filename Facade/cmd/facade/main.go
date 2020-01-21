package main

import (
	"testing"

	"github.com/RyzhAlexWork/go-intern/Facade/pkg/facade"
)

func TestFacade(t *testing.T) {
	expect := "Wallet was create.\nStatus: Nice day! :)\nPay was success."

	user := facade.NewUser()
	result := user.FirstSignIn()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
