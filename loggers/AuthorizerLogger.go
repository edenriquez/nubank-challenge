package loggers

import (
	"log"

	nubankModels "github.com/edenriquez/nubank-challenge/models"
)

// LogError will help us encapsulatin logging error logic
func LogError(errorList ...error) {
	if len(errorList) > 0 {
		log.Println("==== Errors log start ====")
		printErrorList(errorList...)
		log.Fatal("==== Errors log end ====")
	}
}

func printErrorList(errorList ...error) {
	for _, err := range errorList {
		log.Println(err)
	}
}

// LogAction should log Transaction process or account processes
func LogAction(account nubankModels.Account, logType string) {
	out := &nubankModels.OutputJSON{
		Account: nubankModels.Account{
			AccountDetails: account.AccountDetails,
			Violations:     account.Violations,
		},
	}
	log.Print(logType, "\t-\t", out.ToString())
}
