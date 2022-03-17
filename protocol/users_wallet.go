package protocol



// user's wallet
// when we are talking about user's wallet , we are talking about offline transactions
//financial transaction
type Bank_Account_Content struct {
	ID uint64 `json:"id"`
	// whose this wallet is this ?
	UserID uint64 `json:"user_id"`
	// بدهکار
	Debtor string `json:"debtor"`
	// بستانکار
	Creditor string `json:"creditor"`
	// مانده حساب
	Remainder string `json:"remainder"`
	// state  ?
	State string `json:"state"`
	// برداشت پول از بانک
	Withdrawal string `json:"withdrawal"`
	// واریز پول در بانک
	Deposit string `json:"deposit"`
	// انتقال
	// Transition string    `json:"transition"`
}

// how much monety has been tranfereedd
// amount 
// mande
// balance

// financial bank account 


// بدهکار  debtor
// Creditor بستانکار

type Bank_Account interface {
	RegisterWallet() (Bank_Account_Content, error)
	Getwalletremainder(walletID string) (Bank_Account_Content, error)
	GetWalletState()
}
