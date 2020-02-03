package services

import (
	"testing"
	"time"

	"github.com/edenriquez/nubank-challenge/config/constants"

	nubankModel "github.com/edenriquez/nubank-challenge/models"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizeTransactions(t *testing.T) {
	availableLimitBefore := 100
	testAccount := &nubankModel.Account{}
	testAccount.Mock()

	tr1.Mock(20, "test 1", time.Now().UTC().String())
	tr2.Mock(80, "test 2", time.Now().UTC().String())

	testAccount.AvailableLimit = availableLimitBefore
	testAccount.Transactions = []nubankModel.Transaction{}

	testAccount.Transactions = append(testAccount.Transactions, tr1)
	testAccount.Transactions = append(testAccount.Transactions, tr2)

	// AuthorizeTransactions(testAccount)

	totalDiscount := availableLimitBefore - (tr1.Amount + tr2.Amount)

	assert.Equal(t, testAccount.AvailableLimit, totalDiscount)
	assert.Equal(t, testAccount.Violations, []string{""})

}

func TestAuthorizeTransactionsWithInsuficientLimit(t *testing.T) {
	availableLimitBefore := 100
	testAccount := &nubankModel.Account{}
	testAccount.Mock()

	tr1.Mock(20, "test 1", time.Now().UTC().String())
	tr2.Mock(90, "test 2", time.Now().UTC().String())
	tr3.Mock(90, "test 3", time.Now().UTC().String())

	testAccount.AvailableLimit = availableLimitBefore
	testAccount.Transactions = []nubankModel.Transaction{}

	testAccount.Transactions = append(testAccount.Transactions, tr1)
	testAccount.Transactions = append(testAccount.Transactions, tr2)
	testAccount.Transactions = append(testAccount.Transactions, tr3)

	// AuthorizeTransactions(testAccount)

	assert.Equal(t, testAccount.Violations[0], constants.InsuficientLimitError)
}


}
