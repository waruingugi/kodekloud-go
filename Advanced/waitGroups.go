// A wait group is a synchronization primitive that
// allows multiple go-routines to wait for each other (Blocks execution in a structured way).
package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(i int, wg *sync.WaitGroup) {
	fmt.Println(i * i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go calculateSquare(i, &wg)
	}

	elapsed := time.Since(start)
	wg.Wait()
	fmt.Println("Function took", elapsed)
}
