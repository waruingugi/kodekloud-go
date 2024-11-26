package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("delete.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println(err)
		return
	}

	// log.SetOutput(file)
	log.SetOutput(file)
	log.Println("Hello World")
}
