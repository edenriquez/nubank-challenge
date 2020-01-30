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

// Operations represents the result obejct after transaction processing
type Operations struct {
	Violation `json:"violations"`
}

// Account represents the account object
type Account struct {
	ActiveCard     bool `json:"active-card"`
	AvailableLimit int  `json:"available-limit"`
}

// Violation stores string validation based on business logic
type Violation struct {
	Reason string
}

// Transaction represents every transaction associated to the account
type Transaction struct {
	Merchant string `json:"merchant"`
	Amount   int    `json:"amount,int"`
	Time     string `json:"time"`
}
