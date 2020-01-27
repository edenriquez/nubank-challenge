package models

// InputJSON represents the input data structure
type InputJSON struct {
	Account     `json:"account"`
	Transaction `json:"transaction"`
}

// AccountList represents multiple input data
type AccountList struct {
	Account []Account
}

// Account represents the account object
type Account struct {
	ActiveCard     bool          `json:"active-card"`
	AvailableLimit int           `json:"available-limit"`
	Transaction    []Transaction `json:"transaction"`
}

// Transaction represents every transaction associated to the account
type Transaction struct {
	Merchant string `json:"merchant"`
	Amount   int    `json:"amount,int"`
	Time     string `json:"time"`
}
