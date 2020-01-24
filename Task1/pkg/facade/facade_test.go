package facade

import (
	"testing"

	walletstatusPackage "github.com/RyzhAlexWork/go-intern/Task1/pkg/status"
	walletPackage "github.com/RyzhAlexWork/go-intern/Task1/pkg/wallet"
)

var (
	addToBalance   = []int{1000, 15000, 500000, 1000000}
	payFromBalance = []int{5000, 300000, 1500000}
	expectsAdd     = []int{1000, 16000, 516000, 516000}
	expectsPay     = []int{511000, 211000, 211000}
	balance        int
	newWallet      = walletPackage.NewWallet(0, 1000000)
	newStatus      = walletstatusPackage.NewWalletStatus()
)

func Test_Facade(t *testing.T) {
	newUser := NewUser("Valera", newWallet, newStatus)
	for i, expect := range expectsAdd {
		newUser.Add(addToBalance[i])
		balance = newUser.Balance()
		if balance != expect {
			t.Errorf("Wrong balance after deposit! Expect: %d, balance: %d", expect, balance)
		}
	}
	for i, expect := range expectsPay {
		newUser.Pay(payFromBalance[i])
		balance = newUser.Balance()
		if balance != expect {
			t.Errorf("Wrong balance after pay! Expect: %d, balance: %d", expect, balance)
		}
	}
}
