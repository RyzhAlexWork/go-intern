// Package facade is an example of the Facade Pattern.
package facade

import "strings"

// Implements user
type user struct {
	wallet *wallet
	status *status
	pay    *pay
}

// Wallet ...
type wallet struct {
	money int
}

// Status ...
type status struct {
	text string
}

// Pay ...
type pay struct {
	sum int
}

// Create wallet and put in 15000 y.e.
func (w *wallet) create() string {
	w.money = 15000
	return "Wallet was create."
}

// Install auto-status
func (s *status) getStatus() string {
	s.text = "Nice day! :)"
	return "Status: " + s.text
}

// Initial first pay
func (p *pay) payToSite(w *wallet) string {
	p.sum = 1000
	w.money = w.money - p.sum
	return "Pay was success."
}

// FirstSignIn returns user's data after first sign in.
func (u *user) FirstSignIn() string {
	result := []string{
		u.wallet.create(),
		u.status.getStatus(),
		u.pay.payToSite(u.wallet),
	}
	return strings.Join(result, "\n")
}

// NewUser creates user.
func NewUser() *user {
	return &user{
		wallet: &wallet{},
		status: &status{},
		pay:    &pay{},
	}
}
