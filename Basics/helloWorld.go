// There are 3 ways to declare variables and every variable must have a data type associated with it.
// Method 1: var number int = 10
// Method 2: var number = 10
// Method 3: number := 10

package main

import "fmt"

func main() {
	// explicitly declare the data type
	var number1 int = 10
	fmt.Println(number1)

	// assign a value without declaring the data type
	var number2 = 20
	fmt.Println(number2)

	// shorthand notation to define a variable
	number3 := 30
	fmt.Println(number3)

	// initial value
	number := 10
	fmt.Println("Initial number value", number) // prints 10

	// change variable value
	number = 100
	fmt.Println("The changed value", number) // prints 100

	// Multiple variable assignment in a single line
	// var name, age = "Palistha", 22
	// name, age := "Palistha", 22

	// Constants
	// Constatns are fixed values that cannot be changed once declared
	// You cannot use the shorthand notation := to create constants
	const lightSpeed = 299792458
	// const lightSpeed := 299792458 // Results in an error
}
