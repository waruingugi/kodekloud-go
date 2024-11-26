package main

import "fmt"

// func addNumbers() {
// 	n1 := 12
// 	n2 := 8

// 	sum := n1 + n2
// 	fmt.Println("Sum:", sum)
// }

// define a function with 2 parameters
// func addNumbers(n1 int, n2 int) {
// 	sum := n1 + n2
// 	fmt.Println("Sum:", sum)

// 	// return sum
// }

func calculate(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	difference := n1 - n2

	return sum, difference
}

func main() {
	// addNumbers(21, 13)
	sum, difference := calculate(21, 13)
	fmt.Println("Sum:", sum, "Difference: ", difference)
}
