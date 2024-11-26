/*
Channels are a means through which different go routines communicate
Communication in a channel is bi-directional by default

Unbuffered channel: is a channel that needs a receiver as soon as a message is emmitted to the channel.
It has no capacity to hold data.

Buffered channel: has capacity to hold data
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int, 3)
	go sell(ch, &wg)
	wg.Wait()

}

func sell(ch chan int, s *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12

	go buy(ch, s)

	fmt.Println("Sent all data to the channel")
	s.Done()
}

func buy(ch chan int, s *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	fmt.Println("Receiving data: ", <-ch)
	s.Done()
}
