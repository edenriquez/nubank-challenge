package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessStreamToEntity(t *testing.T) {
	objectWithNormalBehaviour := []byte(`
		{"account": {"active-card": true, "available-limit": 100}}
		{"transaction": {"merchant": "Burger King", "amount": 20, "time":"2019-02-13T10:00:00.000Z"}}
		{"transaction": {"merchant": "Habbib's", "amount": 90, "time": "2019-02-13T11:00:00.000Z"}}`)

	acc, err := ProcessStreamToEntity(objectWithNormalBehaviour)
	assert.NotNil(t, acc.AccountDetails.ActiveCard)
	assert.NotNil(t, acc.AccountDetails.AvailableLimit)
	assert.Len(t, acc.Transactions, 2)
	assert.Len(t, err, 0)
}

func TestProcessStreamToEntityWithMalformedInput(t *testing.T) {
	objectWithInvalidStructure := []byte(`
		{"account": {"active-card": true, "available-limit": 100}}
		{"transaction": {"merchant": "Burger King", "amount": 20, "time":"2019-02-13T10:00:00.000Z"}}
		{"transaction": {"merchant": "Habbib's", "amount": 90, "time": "2019-02-13T11:00:00.000Z"}}
		}`)

	_, err := ProcessStreamToEntity(objectWithInvalidStructure)
	assert.Len(t, err, 1)
	assert.Error(t, err[0], `invalid character '}' looking for beginning of value`)
}
