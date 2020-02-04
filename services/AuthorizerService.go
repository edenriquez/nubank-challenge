package services

import (
	"github.com/edenriquez/nubank-challenge/config/constants"
	"github.com/edenriquez/nubank-challenge/loggers"
	nubankLogger "github.com/edenriquez/nubank-challenge/loggers"
	nubankModels "github.com/edenriquez/nubank-challenge/models"
	"github.com/edenriquez/nubank-challenge/utils"
)

// ProcessStreamToEntity should handle incoming lines
func ProcessStreamToEntity(lines []string) (nubankModels.Account, []error) {
	var parseErrorList []error
	account := nubankModels.Account{}
	for _, line := range lines {
		currentLine := &nubankModels.InputJSON{}
		hasError := currentLine.ToStruct(line)
		utils.AppendError(hasError, &parseErrorList)

		if currentLine.AccountIsValid() {
			if account.IsAlreadyCreated() {
				account.AppendValidation(constants.AccountAlreadyInitialized)
				nubankLogger.LogAction(account, "creation")
				break
			}
			account = currentLine.Account
			nubankLogger.LogAction(account, "creation")
		} else if currentLine.IsTransaction() {
			account.Transactions = append(account.Transactions, currentLine.Transaction)
			if !account.IsAlreadyCreated() {
				account.AppendValidation(constants.AccountIsNotInitialized)
				nubankLogger.LogAction(account, "transaction")
				break
			}
		} else if !currentLine.AccountIsValid() {
			account.AppendValidation(constants.AccountCardIsNotActive)
			nubankLogger.LogAction(account, "creation")
			break
		}
	}
	return account, parseErrorList
}

// ProcessEntityTransactions should validate and charge transactions on account
func ProcessEntityTransactions(account *nubankModels.Account) {
	if len(account.Violations) == 0 {
		for _, transaction := range account.Transactions {
			violation := account.ChargeTransaction(transaction)
			account.AppendValidation(violation)
			loggers.LogAction(*account, "transaction")
		}
	}
}
