package main

import (
	"io/ioutil"
	"os"

	nubankControllers "github.com/edenriquez/nubank-challenge/controllers"
	nubankLogger "github.com/edenriquez/nubank-challenge/loggers"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	account, errorList := nubankControllers.ProcessStreamToEntity(data)
	nubankLogger.LogError(errorList...)
	nubankControllers.ProcessEntityTransactions(&account)
}
