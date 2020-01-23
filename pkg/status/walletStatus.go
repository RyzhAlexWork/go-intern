package walletstatus

// WalletStatus ...
type WalletStatus interface {
	Get() (text string)
	Change(newText string)
}

type walletStatus struct {
	text string
}

// Get return walletStatus text
func (w *walletStatus) Get() (text string) {
	return w.text
}

// Change change the walletStatus text
func (w *walletStatus) Change(newText string) {
	w.text = newText
}

// NewWalletStatus create status implementation for interface WalletStatus
func NewWalletStatus() WalletStatus {
	return &walletStatus{}
}
