package models

import (
	"encoding/json"
	"time"

	"github.com/edenriquez/nubank-challenge/config/constants"

	"gopkg.in/validator.v2"
)

const layout = "2006-01-02T15:04:05.000Z"

// ChargeTransaction method that validates and updates account available limit
func (a *Account) ChargeTransaction(transaction Transaction) string {
	updatedAccount := &Account{
		AccountDetails: a.AccountDetails,
	}
	updatedAccount.AvailableLimit -= transaction.Amount
	if err := validator.Validate(updatedAccount); err != nil {
		return constants.InsuficientLimitError
	}
	if a.hasDuplicateTransaction(transaction) {
		return constants.DoubleTransaction
	}
	if a.hasHighFrequency(transaction) {
		return constants.TransactionHighFrequency
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

func (a *Account) hasHighFrequency(incomingTransaction Transaction) bool {
	highFrequencyCounter := 1
	for _, accountTransaction := range a.Transactions {
		accountTime, _ := time.Parse(layout, accountTransaction.Time)
		incomingTime, _ := time.Parse(layout, incomingTransaction.Time)
		twoMinAgo := accountTime.Add(time.Duration(-2) * time.Minute)
		twoMinAfter := accountTime.Add(time.Duration(2) * time.Minute)
		isInBetweenRange := incomingTime.Before(twoMinAfter) && incomingTime.After(twoMinAgo)
		if isInBetweenRange {
			highFrequencyCounter++
		}
	}
	return highFrequencyCounter > 3
}

func (a *Account) hasDuplicateTransaction(incomingTransaction Transaction) bool {
	duplicate := false
	for _, accountTransaction := range a.Transactions {
		hasSameMerchant := accountTransaction.Merchant == incomingTransaction.Merchant
		hasSameAmount := accountTransaction.Amount == incomingTransaction.Amount
		differentTime := accountTransaction.Time != incomingTransaction.Time
		if hasSameAmount && hasSameMerchant && differentTime {
			accountTime, _ := time.Parse(layout, accountTransaction.Time)
			incomingTime, _ := time.Parse(layout, incomingTransaction.Time)
			twoMinAgo := accountTime.Add(time.Duration(-2) * time.Minute)
			twoMinAfter := accountTime.Add(time.Duration(2) * time.Minute)
			isInBetweenRange := incomingTime.Before(twoMinAfter) && incomingTime.After(twoMinAgo)
			if isInBetweenRange {
				duplicate = true
			}
		}
	}
	return duplicate
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
