package services

import (
	"testing"
	"time"

	"github.com/edenriquez/nubank-challenge/config/constants"

	nubankModel "github.com/edenriquez/nubank-challenge/models"
	"github.com/stretchr/testify/assert"
)

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
