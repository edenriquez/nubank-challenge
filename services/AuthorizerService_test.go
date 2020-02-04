package services

import (
	"testing"
	"time"

	"github.com/edenriquez/nubank-challenge/config/constants"

	nubankModel "github.com/edenriquez/nubank-challenge/models"
	"github.com/stretchr/testify/assert"
)

func TestProcessStreamToEntity(t *testing.T) {
	inputString := []string{
		`{"account": {"active-card": true, "available-limit": 100}}`,
		`{"transaction": {"merchant": "Burger King", "amount": 20, "time":"2019-02-13T10:00:00.000Z"}}`,
		`{"transaction": {"merchant": "Habbib's", "amount": 90, "time": "2019-02-13T11:00:00.000Z"}}`,
	}

	account, err := ProcessStreamToEntity(inputString)
	assert.Nil(t, err)
	assert.Equal(t, account.AccountDetails.AvailableLimit, 100)
	assert.True(t, account.AccountDetails.ActiveCard)
	assert.Len(t, account.Violations, 0)
	assert.Len(t, account.Transactions, 2)
}

func TestProcessStreamToEntityWithTwoAccounts(t *testing.T) {
	inputString := []string{
		`{"account": {"active-card": true, "available-limit": 100}}`,
		`{"transaction": {"merchant": "Burger King", "amount": 20, "time":"2019-02-13T10:00:00.000Z"}}`,
		`{"transaction": {"merchant": "Habbib's", "amount": 90, "time": "2019-02-13T11:00:00.000Z"}}`,
		`{"account": {"active-card": true, "available-limit": 100}}`,
	}

	account, err := ProcessStreamToEntity(inputString)
	assert.Nil(t, err)
	assert.Equal(t, account.AccountDetails.AvailableLimit, 100)
	assert.True(t, account.AccountDetails.ActiveCard)
	assert.Len(t, account.Transactions, 2)
	assert.Len(t, account.Violations, 1)
	assert.Equal(t, account.Violations[0], constants.AccountAlreadyInitialized)
}

func TestProcessStreamToEntityWithoutAccount(t *testing.T) {
	inputString := []string{
		`{"transaction": {"merchant": "Burger King", "amount": 20, "time":"2019-02-13T10:00:00.000Z"}}`,
		`{"account": {"active-card": true, "available-limit": 100}}`,
		`{"transaction": {"merchant": "Habbib's", "amount": 90, "time": "2019-02-13T11:00:00.000Z"}}`,
	}
	account, err := ProcessStreamToEntity(inputString)

	assert.Nil(t, err)
	assert.Len(t, account.Violations, 1)
	assert.Equal(t, account.Violations[0], constants.AccountIsNotInitialized)
}

func TestProcessStreamToEntityWithCardNotActive(t *testing.T) {
	inputString := []string{
		`{"account": {"active-card": false, "available-limit": 100}}`,
		`{"transaction": {"merchant": "Burger King", "amount": 20, "time":"2019-02-13T10:00:00.000Z"}}`,
		`{"transaction": {"merchant": "Habbib's", "amount": 90, "time": "2019-02-13T11:00:00.000Z"}}`,
	}
	account, err := ProcessStreamToEntity(inputString)

	assert.Nil(t, err)
	assert.Len(t, account.Violations, 1)
	assert.Equal(t, account.Violations[0], constants.AccountCardIsNotActive)
}

func TestProcessStreamToEntityWithDuplicateTransaction(t *testing.T) {
	inputString := []string{
		`{"account": {"active-card": true, "available-limit": 1000}}`,
		`{"transaction": {"merchant": "Habbib's", "amount": 1, "time": "2019-02-13T11:00:00.000Z"}}`,
		`{"transaction": {"merchant": "Habbib's", "amount": 1, "time": "2019-02-13T11:01:00.000Z"}}`,
		`{"transaction": {"merchant": "Habbib's", "amount": 1, "time": "2019-02-13T11:02:00.000Z"}}`,
	}
	account, err := ProcessStreamToEntity(inputString)
	ProcessEntityTransactions(&account)
	assert.Nil(t, err)
	assert.Len(t, account.Violations, 1)
	assert.Equal(t, account.Violations[0], constants.DoubleTransaction)

}

func TestProcessStreamToEntityWithTransactionHighFrequency(t *testing.T) {
	inputString := []string{
		`{"account": {"active-card": true, "available-limit": 1000}}`,
		`{"transaction": {"merchant": "fraud 1", "amount": 20, "time":"2019-02-13T10:00:00.000Z"}}`,
		`{"transaction": {"merchant": "fraud 2", "amount": 10, "time": "2019-02-13T10:01:00.000Z"}}`,
		`{"transaction": {"merchant": "fraud 3", "amount": 1, "time": "2019-02-13T10:01:01.000Z"}}`,
	}
	account, err := ProcessStreamToEntity(inputString)
	ProcessEntityTransactions(&account)
	assert.Nil(t, err)
	assert.Len(t, account.Violations, 1)
	assert.Equal(t, account.Violations[0], constants.TransactionHighFrequency)

}

func TestProcessEntityTransactions(t *testing.T) {
	availableLimitBefore := 100
	testAccount := &nubankModel.Account{}
	testAccount.Mock()
	tr1 := nubankModel.Transaction{}
	tr2 := nubankModel.Transaction{}
	tr1.Mock(20, "test 1", time.Now().UTC().String())
	tr2.Mock(80, "test 2", time.Now().UTC().String())

	testAccount.AvailableLimit = availableLimitBefore
	testAccount.Transactions = []nubankModel.Transaction{}

	testAccount.Transactions = append(testAccount.Transactions, tr1)
	testAccount.Transactions = append(testAccount.Transactions, tr2)

	ProcessEntityTransactions(testAccount)

	totalDiscount := availableLimitBefore - (tr1.Amount + tr2.Amount)

	assert.Equal(t, testAccount.AvailableLimit, totalDiscount)
	assert.Len(t, testAccount.Violations, 0)

}

func TestProcessEntityTransactionsWithInsuficientLimitError(t *testing.T) {
	availableLimitBefore := 80
	testAccount := &nubankModel.Account{}
	testAccount.Mock()
	tr1 := nubankModel.Transaction{}
	tr2 := nubankModel.Transaction{}
	tr1.Mock(20, "test 1", time.Now().UTC().String())
	tr2.Mock(80, "test 2", time.Now().UTC().String())

	testAccount.AvailableLimit = availableLimitBefore
	testAccount.Transactions = []nubankModel.Transaction{}

	testAccount.Transactions = append(testAccount.Transactions, tr1)
	testAccount.Transactions = append(testAccount.Transactions, tr2)

	ProcessEntityTransactions(testAccount)

	assert.Equal(t, testAccount.AvailableLimit, availableLimitBefore-tr1.Amount)
	assert.Len(t, testAccount.Violations, 1)
	assert.Equal(t, testAccount.Violations[0], constants.InsuficientLimitError)

}
