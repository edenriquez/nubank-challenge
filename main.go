package main

import (
	"fmt"
	"io/ioutil"
	"os"

	nubankControllers "github.com/edenriquez/nubank-challenge/controllers"
	nubankLogger "github.com/edenriquez/nubank-challenge/loggers"
	nubankModel "github.com/edenriquez/nubank-challenge/models"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	accounts, transactions, errorList := nubankControllers.GenerateModel(data)
	nubankLogger.LogError(errorList...)
	operations := nubankControllers.AuthorizeTransactions(accounts, transactions)
	printOperations(operations...)
}

func printOperations(o ...nubankModel.Operations) {
	for _, operation := range o {
		fmt.Println(operation)
	}
}
