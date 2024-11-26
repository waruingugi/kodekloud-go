/*
Switch: Non blocking.
They are also deterministic and will run in sequence to select the matching case.
Select: statements can block since they are used with channels, and they can block or receive operation.
Are non-deterministic, as it will execute a case randomly with no sequence.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	default:
		fmt.Println("Executed default block")
	}

	time.Sleep(1 * time.Second)
}

func goTwo(ch2 chan string) {
	ch2 <- "Channel - 2"
}

func goOne(ch1 chan string) {
	ch1 <- "Channel - 1"
}
