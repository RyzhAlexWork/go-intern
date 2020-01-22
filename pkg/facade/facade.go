// Package facade is an example of the Facade Pattern.
package facade

import (
	"strings"
)
// Implements user
type User interface {
	FirstSignIn() string
}

type walleter interface {
	Create() string
}

type statuser interface {
	GetStatus()	string
}

type payer interface {
	PayToSite() string
}


type user struct {
	wallet walleter
	status statuser
	pay    payer
}

// FirstSignIn returns user's data after first sign in.
func (u *user) FirstSignIn() string {
	result := []string{
		u.wallet.Create(),
		u.status.GetStatus(),
		u.pay.PayToSite(),
	}
	return strings.Join(result, "\n")
}

// NewUser creates user.
func NewUser(inputWallet walleter, inputStatus statuser, inputPay payer) User {
	return &user{
		wallet: inputWallet,
		status: inputStatus,
		pay:    inputPay,
	}
}
