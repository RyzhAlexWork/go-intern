// Package facade is an example of the Facade Pattern.
package facade

import "strings"

// Implements user
type User struct {
	wallet *Wallet
	status *Status
	pay    *Pay
}

// Wallet ...
type Wallet struct {
	money int
}

// Status ...
type Status struct {
	text string
}

// Pay ...
type Pay struct {
	sum int
}

// Create wallet and put in 15000 y.e.
func (w *Wallet) Create() string {
	w.money = 15000
	return "Wallet was create."
}

// Install auto-status
func (s *Status) GetStatus() string {
	s.text = "Nice day! :)"
	return "Status: " + s.text
}

// Initial first pay
func (p *Pay) PayToSite(w *Wallet) string {
	p.sum = 1000
	w.money = w.money - p.sum
	return "Pay was success."
}

// FirstSignIn returns user's data after first sign in.
func (u *User) FirstSignIn() string {
	result := []string{
		u.wallet.Create(),
		u.status.GetStatus(),
		u.pay.PayToSite(u.wallet),
	}
	return strings.Join(result, "\n")
}

// NewUser creates user.
func NewUser() *User {
	return &User{
		wallet: &Wallet{},
		status: &Status{},
		pay:    &Pay{},
	}
}
