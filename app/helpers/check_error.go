package helpers

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		panic(err)
	}
}
