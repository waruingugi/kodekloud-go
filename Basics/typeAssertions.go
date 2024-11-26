package main

import "fmt"

func main() {
	// create an empty interface
	var a interface{}

	// store integer to an empty interface
	a = 10

	// type assertion
	interfaceValue, booleanValue := a.(int)
	fmt.Println("Interface Value:", interfaceValue)
	fmt.Println("Boolean Value:", booleanValue)
}
