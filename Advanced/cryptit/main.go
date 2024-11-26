package main

import (
	"cryptit/decrypt"
	"cryptit/encrypt"
	"fmt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
