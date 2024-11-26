package main

import "fmt"

func main() {
	// number := 15

	// true if number if less than 0
	// if number > 0 {
	// 	fmt.Printf("%d is a positive number\n", number)
	// }
	// fmt.Println("Out of the loop")

	// number := 10
	// if number > 0 {
	// 	fmt.Printf("%d is a positive number\n", number)
	// } else {
	// 	fmt.Printf("%d is a negative number\n", number)
	// }

	// number1 := 12
	// number2 := 20

	// if number1 == number2 {
	// 	fmt.Printf("Result: %d == %d", number1, number2)
	// } else if number1 > number2 {
	// 	fmt.Printf("Result: %d > %d", number1, number2)
	// } else {
	// 	fmt.Printf("Result: %d < %d", number1, number2)
	// }

	number1 := 12
	number2 := 20

	// Outer if statement
	if number1 >= number2 {
		if number1 == number2 {
			fmt.Printf("Result: %d == %d", number1, number2)
			// inner else statement
		} else {
			fmt.Printf("Result: %d > %d", number1, number2)
		}
	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}
}
