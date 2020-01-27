package main

import (
	"fmt"
	"io/ioutil"
	"os"

	nubankControllers "github.com/edenriquez/nubank-authorizer/controllers"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	accounts, err := nubankControllers.GenerateModel(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(accounts.Account))
	fmt.Println(len(accounts.Account[0].Transaction))

}
