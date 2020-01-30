package loggers

import (
	"log"
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
