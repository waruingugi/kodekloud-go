package main

import "fmt"

// empty interface as function parameter
// func displayValue(i interface{}) {
// 	fmt.Println(i)
// }

// function with an empty interface as its parameter
func displayValue(i ...interface{}) {
	fmt.Println(i)
}

func main() {
	a := "Welcome to Programiz"
	b := 20
	c := false

	// // pass string to the function
	// displayValue(a)
	displayValue(a)

	// // pass integer number to the function
	// displayValue(b)
	displayValue(a, b)

	// // pass boolean value to the function
	// displayValue(c)
	displayValue(a, b, c)
}
