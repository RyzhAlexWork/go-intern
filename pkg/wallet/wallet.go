package wallet

// Wallet ...
type Wallet interface {
	Create() string
	Pay(sum int) bool
}

type wallet struct {
	money int
}

// Create wallet and put in 15000 y.e.
func (w *wallet) Create() string {
	w.money = 15000
	return "Wallet was create."
}

func (w *wallet) Pay(sum int) bool {
	w.money = w.money - sum
	return true
}

func NewWallet() Wallet {
	return  &wallet{}
}

