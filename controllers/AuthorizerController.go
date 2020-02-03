package controllers

import (
	nubankModels "github.com/edenriquez/nubank-challenge/models"
	"github.com/edenriquez/nubank-challenge/services"
	"github.com/edenriquez/nubank-challenge/utils"
)

// ProcessStreamToEntity cast inconming json string into AccountList struct
func ProcessStreamToEntity(data []byte) (nubankModels.Account, []error) {
	lines := utils.ParseByteToStringLines(data)
	return services.ProcessStreamToEntity(lines)
}

// ProcessEntityTransactions should trigger account process to validate transactions
func ProcessEntityTransactions(account *nubankModels.Account) {
	services.ProcessEntityTransactions(account)
}
