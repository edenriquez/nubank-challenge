package models

// InputJSON represents the input data structure
type InputJSON struct {
	Account     `json:"account"`
	Transaction `json:"transaction"`
}

// OutputJSON represents the ouput data structure
// with Account and Violations objects
// Transactions is not required in the output so is avoided with - tag
type OutputJSON struct {
	Account `json:"account"`
}

// Account represents in memory structure for given json
type Account struct {
	AccountDetails
	Transactions []Transaction `json:"-"` // - ommit this field when parsing bytes to json
	Violations   []string      `json:"violations,omitempty"`
}

// AccountDetails represents the account object
type AccountDetails struct {
	ActiveCard     bool `json:"active-card" validate:"nonnil"`
	AvailableLimit int  `json:"available-limit" validate:"nonnil,min=0"`
}

// Transaction represents every transaction associated to the account
type Transaction struct {
	Merchant string `json:"merchant"`
	Amount   int    `json:"amount,int"`
	Time     string `json:"time"`
}

// Violation stores string validation based on business logic
type Violation struct {
	string
}
