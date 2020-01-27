package controller

import (
	"encoding/json"
	"strings"

	nubankModel "github.com/edenriquez/nubank-authorizer/models"
)

// GenerateModel cast inconming json string into AccountList struct
func GenerateModel(data []byte) (nubankModel.AccountList, error) {
	var err error
	var lines []string
	lastIndex := -1
	lines = parseLines(data)
	accountList := nubankModel.AccountList{}
	for _, line := range lines {
		currentLine := nubankModel.InputJSON{}
		err = json.Unmarshal([]byte(line), &currentLine)
		if currentLine.Account.ActiveCard {
			accountList.Account = append(accountList.Account, currentLine.Account)
			lastIndex++
		} else if len(currentLine.Transaction.Merchant) > 0 {
			accountList.Account[lastIndex].Transaction = append(
				accountList.Account[lastIndex].Transaction,
				currentLine.Transaction)
		}
	}
	return accountList, err
}

func parseLines(data []byte) []string {
	lines := strings.Split(string(data), "\n")
	var result []string
	for _, line := range lines {
		if len(line) > 0 {
			result = append(result, line)
		}
	}
	return result
}
