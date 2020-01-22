package pay

type wallet interface {
	Pay(sum int) bool
}

// Pay ...
type Pay interface {
	PayToSite() string
}

type pay struct {
	w	wallet
	sum int
}

// Initial first pay
func (p *pay) PayToSite() string {
	if p.w.Pay(p.sum) {
		return "Pay was success."
	} else {
		return "Pay was crash"
	}
}

func NewPay(inputWallet wallet, sum int) Pay {
	return &pay{
		sum:	sum,
		w:		inputWallet,
	}
}
