package facade

import (
	"testing"

	"github.com/RyzhAlexWork/go-intern/pkg/pay"
	"github.com/RyzhAlexWork/go-intern/pkg/status"
	"github.com/RyzhAlexWork/go-intern/pkg/wallet"
)

var (
	expect = "Wallet was create.\nStatus: Nice day! :)\nPay was success."
	sums = []int{1000, 500, 1500, 2500, 6000}
)

func Test_Facade(t *testing.T) {
	for sum := range sums {
		wallet := wallet.NewWallet()
		status := status.NewStatus()
		pay := pay.NewPay(wallet,sum)
		user := NewUser(wallet, status, pay)
		result := user.FirstSignIn()
		if result != expect {
			t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
		}
	}
}
