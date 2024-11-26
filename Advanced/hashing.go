package main

import (
	"crytpo/md5"
	"encoding/hex"
	"fmt"
)

func encodeWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])
}

func main() {
	var password string
	fmt.Scanln(&password)
	fmt.Println("Password encrypted to:", encodeWithMD(password))
}
