package main

import (
	"fmt"

	"github.com/pborman/uuid"

	"cryptit/encrypt"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
}
