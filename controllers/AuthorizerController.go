package controllers

import (
	"encoding/json"
	"strings"

	constants "github.com/edenriquez/nubank-challenge/config/constants"
	nubankModel "github.com/edenriquez/nubank-challenge/models"
)

var lines []string
var globalAccount nubankModel.Account
var operations []nubankModel.Operations
var account []nubankModel.Account
var transactions []nubankModel.Transaction
var listOfErrors []error

// GenerateModel cast inconming json string into AccountList struct
func GenerateModel(data []byte) ([]nubankModel.Account, []nubankModel.Transaction, []error) {
	lines = parseLines(data)
	for _, line := range lines {
		currentLine := nubankModel.InputJSON{}
		appendError(json.Unmarshal([]byte(line), &currentLine))
		if currentLine.Account.ActiveCard {
			account = append(account, currentLine.Account)
		} else if len(currentLine.Transaction.Merchant) > 0 {
			transactions = append(transactions, currentLine.Transaction)
		}
	}
	return account, transactions, listOfErrors
}

// AuthorizeTransactions should validate transactions for given account
func AuthorizeTransactions(accounts []nubankModel.Account, transactions []nubankModel.Transaction) []nubankModel.Operations {
	globalAccount = getAssociatedAccount(accounts...)
	hasConstraints := false
	for _, transaction := range transactions {
		violation := chargeTransaction(transaction)
		if len(violation.Reason) > 0 {
			operations = append(operations, nubankModel.Operations{
				Violation: violation,
			})
			hasConstraints = true
		}
	}
	if hasConstraints {
		return operations
	}
	return nil
}

func chargeTransaction(transaction nubankModel.Transaction) nubankModel.Violation {
	if globalAccount.AvailableLimit-transaction.Amount > 0 {
		globalAccount.AvailableLimit -= transaction.Amount
	} else {
		// TODO implement dynamic assignation here
		return nubankModel.Violation{
			Reason: constants.InsuficientLimitError,
		}
	}
	return nubankModel.Violation{}
}

func getAssociatedAccount(accounts ...nubankModel.Account) nubankModel.Account {
	return accounts[0]
}

func parseLines(data []byte) []string {
	linesByRow := strings.Split(string(data), "\n")
	var linesWithoutEOF []string
	for _, row := range linesByRow {
		if len(row) > 0 {
			linesWithoutEOF = append(linesWithoutEOF, row)
		}
	}
	return linesWithoutEOF
}

func appendError(e error) {
	if e != nil {
		listOfErrors = append(listOfErrors, e)
	}
}
