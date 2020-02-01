package controllers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCases struct {
	input []byte
}

func TestGenerateModel(t *testing.T) {
	testNormalBehaviour := &testCases{}
	testNormalBehaviour.input = []byte(`{"account": {"active-card": true, "available-limit": 100}}
	{"transaction": {"merchant": "Burger King", "amount": 20, "time":"2019-02-13T10:00:00.000Z"}}
	{"transaction": {"merchant": "Habbib's", "amount": 90, "time": "2019-02-13T11:00:00.000Z"}}`)

	acc, trans, err := GenerateModel(testNormalBehaviour.input)
	fmt.Println(acc, trans)
	assert.Len(t, acc, 1)
	assert.Len(t, trans, 2)
	assert.Len(t, err, 0)

}
