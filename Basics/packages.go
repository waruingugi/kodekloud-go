/*
Print(): prints the text to output screen
Println(): prints the text to output with a new line character at the end
Printf(): prints the formatted string to the output screen

Scan(): get input values from the user
Scanf(): get input values using the format specifier
Scanln(): get input values until the new line is detected

*/

package main

import (
	"calculator"
	"fmt"
)

func main() {
	// var number int

	// // take input value
	// fmt.Scan(&number)

	// fmt.Println("Number is", number)

	// var number int
	// fmt.Scanf("%d", &number)
	// fmt.Printf("%d", number)

	number1 := 9
	number2 := 5

	// us the add function of calculator package
	fmt.Println(calculator.Add(number1, number2))
}
