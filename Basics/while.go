package main

import "fmt"

func main() {
	// number := 1

	// for number <= 5 {
	// 	fmt.Println(number)
	// 	number++
	// }
	number := 1

	for {
		if number > 5 {
			break
		}

		fmt.Printf("%d\n", number)
		number++
	}
}
