package models

import (
	"encoding/json"
	"time"

	"github.com/edenriquez/nubank-challenge/config/constants"

	"gopkg.in/validator.v2"
)

// ChargeTransaction method that validates and updates account available limit
func (a *Account) ChargeTransaction(transaction Transaction) string {
	updatedAccount := &Account{
		AccountDetails: a.AccountDetails,
	}
	updatedAccount.AvailableLimit -= transaction.Amount
	if err := validator.Validate(updatedAccount); err != nil {
		return constants.InsuficientLimitError
	}
	a.AccountDetails = updatedAccount.AccountDetails
	return ""
}

// AppendValidation helper to add given validation to the account struct
func (a *Account) AppendValidation(v string) {
	notExist := true
	for _, violation := range a.Violations {
		if violation == v {
			notExist = false
		}
	}
	if notExist && len(v) > 0 {
		a.Violations = append(a.Violations, v)
	}
}

// ToString should cast output struct to json formatted
func (o *OutputJSON) ToString() string {
	result, _ := json.Marshal(o)
	return string(result)
}

// ToStruct should cast input string line to struct
func (i *InputJSON) ToStruct(line string) error {
	return json.Unmarshal([]byte(line), &i)
}

// AccountIsValid should help validating if account is already created
func (i *InputJSON) AccountIsValid() bool {
	return i.Account.ActiveCard
}

// IsTransaction should help validating if account is already created
func (i *InputJSON) IsTransaction() bool {
	return len(i.Transaction.Merchant) > 0
}

// IsAlreadyCreated is a helper to check is active card is already initialized
func (a *Account) IsAlreadyCreated() bool {
	return a.ActiveCard
}

// Mock should serve as factory method for Account with dummy data for testing purposes
func (t *Transaction) Mock(amount int, merchant string, time string) {
	t.Amount = amount
	t.Merchant = merchant
	t.Time = time
}

// Mock should serve as factory method for Account with dummy data for testing purposes
func (a *Account) Mock() {
	a.AccountDetails = AccountDetails{
		ActiveCard:     true,
		AvailableLimit: 100,
	}
	a.Transactions = []Transaction{
		{
			Amount:   20,
			Merchant: "Test1",
			Time:     time.Now().UTC().String(),
		},
	}
}
