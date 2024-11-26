// fmt.Print()
// fmt.Println()
// fmt.Printf()
package main

import "fmt"

func main() {
	// Print
	fmt.Print("Hello , ")
	fmt.Print("World!")

	// Print variables
	name := "John"
	fmt.Print(name)
	fmt.Print("Name: ", name)

	// PrintLn
	currentSalary := 50000
	fmt.Println("Current Salary:", currentSalary)

	// Printf
	// Data types are %d (integer), %g (float), %s (string), %t (bool)
	currentAge := 21
	fmt.Printf("Age = %d", currentAge)

	var annualSalary float32 = 65000.5
	fmt.Printf("Annual Salary: %g", annualSalary)

	var userName = "John"
	age := 23
	fmt.Printf("%s is %d years old.", userName, age)

	// if you don't want to use the fmt package, use the below options
	println("Using println instead of fmt.Println")
	print("Using print instead of fmt.Print")
}
