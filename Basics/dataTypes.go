/*
In Go, there are two types of integers:
- Signed int: can hold both positive and negative integers
- Unsigned int: can only hold positive integers
*/
package main

import "fmt"

func main() {
	// Integers
	var integer1 int
	var integer2 int

	integer1 = 5
	integer2 = 10

	fmt.Println(integer1)
	fmt.Println(integer2)

	// float32: 32 bits (4 bytes)
	// float64: 64 bits (8 bytes)
	var salary1 float32
	var salary2 float64

	salary1 = 50000.503882901
	// can store decimals with greater precision
	salary2 = 50000.503882901

	fmt.Println(salary1)
	fmt.Println(salary2)

	// String
	var message string
	message = "Welcome to Programiz"
	fmt.Println(message)

	// Boolean
	var boolValue bool
	boolValue = false
	fmt.Println(boolValue)

}
