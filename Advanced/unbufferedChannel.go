package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sell(ch)
	go buy(ch)
	time.Sleep(2 * time.Second)
}

// sends data to the channel
func sell(ch chan string) {
	ch <- "Furniture"
	fmt.Println("Send data to the channel")
}

// receive data from the channel
func buy(ch chan string) {
	fmt.Println("Waiting for data")
	val := <-ch
	fmt.Println("Received data - ", val)
}
