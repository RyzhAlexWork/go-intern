package facade

// User ...
type User interface {
	Add(money int) (walletStatus string)
	Pay(money int) (walletStatus string)
	Balance() (money int)
}

type wallet interface {
	Pay(amount int) (done bool)
	Add(amount int) (done bool)
	Balance() (walletBalance int)
}

type walletStatus interface {
	Get() (text string)
	Change(newText string)
}

type user struct {
	login        string
	wallet       wallet
	walletStatus walletStatus
}

// Add deposit to wallet and change the walletStatus text
func (u *user) Add(money int) (walletStatus string) {
	var (
		check bool
	)
	check = u.wallet.Add(money)
	if check {
		u.walletStatus.Change("Deposit was successful.")
	} else {
		u.walletStatus.Change("Deposit error. Too much money.")
	}
	return u.walletStatus.Get()
}

// Add makes payment from wallet and change the walletStatus text
func (u *user) Pay(money int) (walletStatus string) {
	var (
		check bool
	)
	check = u.wallet.Pay(money)
	if check {
		u.walletStatus.Change("Pay was successful.")
	} else {
		u.walletStatus.Change("Pay error. Not enough money.")
	}
	return u.walletStatus.Get()
}

// Balance show balance
func (u *user) Balance() (walletBalance int) {
	return u.wallet.Balance()
}

// NewUser creates user.
func NewUser(login string, inputWallet wallet, inputWalletStatus walletStatus) User {
	return &user{
		login:        login,
		wallet:       inputWallet,
		walletStatus: inputWalletStatus,
	}
}
