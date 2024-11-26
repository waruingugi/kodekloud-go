package main

import "fmt"

func main() {
	// var num int = 5
	// // print the value store in variable
	// fmt.Println("Variable value: ", num)

	// // prints the address of the variable
	// fmt.Println("Memory address: ", &num)

	// var name = "John"
	// var ptr *string

	// // assign the memory address of name to the pointer
	// ptr = &name

	// // fmt.Println("Value of pointer is", ptr)
	// // fmt.Println("Address of the variable", &name)
	// fmt.Println(*ptr)

	// var num int
	// var ptr *int

	// num = 22
	// fmt.Println("Addrerss of num:", &num)
	// fmt.Println("Value of num:", num)

	// ptr = &num
	// fmt.Println("\nAddress of pointer ptr:", ptr)
	// fmt.Println("Content of pointer ptr:", *ptr)

	// num = 11
	// fmt.Println("\nAddress of pointer ptr:", ptr)
	// fmt.Println("Content of pointer ptr", *ptr)

	// *ptr = 2
	// fmt.Println("\nAddress of num:", &num)
	// fmt.Println("Value of num:", num)

	// create a pointer using new()
	var ptr = new(int)
	*ptr = 20

	fmt.Println(ptr)
	fmt.Println(*ptr)
}
