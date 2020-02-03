package models

import (
	"testing"
	"time"

	"github.com/edenriquez/nubank-challenge/config/constants"

	"github.com/stretchr/testify/assert"
)

const accountLine = "{\"account\":{\"active-card\":true,\"available-limit\":101}}"

func TestChargeTransaction(t *testing.T) {
	totalAmount := 100
	account := Account{}
	transaction := Transaction{}

	account.Mock()
	account.AccountDetails.AvailableLimit = totalAmount
	transaction.Mock(20, "test 1", time.Now().UTC().String())

	account.Transactions = []Transaction{transaction}
	result := account.ChargeTransaction(transaction)

	assert.Empty(t, result)
	assert.Equal(t, totalAmount-transaction.Amount, account.AccountDetails.AvailableLimit)
}

func TestChargeTransactionInsuficientLimitError(t *testing.T) {
	totalAmount := 10
	account := Account{}
	transaction := Transaction{}

	account.Mock()
	account.AccountDetails.AvailableLimit = totalAmount
	transaction.Mock(20, "test 1", time.Now().UTC().String())

	account.Transactions = []Transaction{transaction}
	result := account.ChargeTransaction(transaction)

	assert.Equal(t, constants.InsuficientLimitError, result)
	assert.Equal(t, totalAmount, account.AccountDetails.AvailableLimit)
}

func TestAppendValidation(t *testing.T) {
	account := Account{}
	account.Mock()

	assert.Nil(t, account.Violations)
	account.AppendValidation(constants.AccountAlreadyInitialized)
	assert.Equal(t, constants.AccountAlreadyInitialized, account.Violations[0])
}

func TestAppendSameValidationTwice(t *testing.T) {
	account := Account{}
	account.Mock()

	assert.Nil(t, account.Violations)
	account.AppendValidation(constants.AccountAlreadyInitialized)
	account.AppendValidation(constants.AccountAlreadyInitialized)

	assert.Len(t, account.Violations, 1)
	assert.Equal(t, constants.AccountAlreadyInitialized, account.Violations[0])
}

func TestOutputJSON(t *testing.T) {
	out := &OutputJSON{
		Account: Account{},
	}
	out.Account.Mock()
	out.Account.AvailableLimit = 101
	assert.IsType(t, out.ToString(), "")
	assert.Equal(t, out.ToString(), accountLine)
}

func TestToStruct(t *testing.T) {
	in := &InputJSON{}
	err := in.ToStruct(accountLine)
	assert.Nil(t, err)
	assert.True(t, in.ActiveCard)
	assert.Equal(t, in.AvailableLimit, 101)
}

func TestToStructError(t *testing.T) {
	in := &InputJSON{}
	err := in.ToStruct(accountLine + "}")
	assert.NotNil(t, err)
}

func TestAccountIsValid(t *testing.T) {
	in := &InputJSON{}
	in.Account.Mock()
	assert.True(t, in.AccountIsValid())
}

func TestAccountIsNotValid(t *testing.T) {
	in := &InputJSON{}
	in.Account.Mock()
	in.Account.ActiveCard = false
	assert.False(t, in.AccountIsValid())
}

func TestIsTransaction(t *testing.T) {
	in := &InputJSON{}
	in.Transaction.Mock(10, "merchant", "01/01/01")
	assert.True(t, in.IsTransaction())
}

func TestIsTransactionFalse(t *testing.T) {
	in := &InputJSON{}
	in.Transaction.Mock(10, "", "01/01/01")
	assert.False(t, in.IsTransaction())
}
