package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {
	// 	// terminates the loop when i is equal to 3
	// 	if i == 3 {
	// 		break
	// 	}

	// 	fmt.Println(i)
	// }

	for i := 1; i <= 5; i++ {
		// skips the iteration when i is equal to 3
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
