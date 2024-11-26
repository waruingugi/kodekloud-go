// fmt.Scan()
// fmt.Scanln()
// fmt.Scanf()

package main

import "fmt"

func main() {
	// Scan() can only take input values up to a space
	var language string
	fmt.Print("Enter your favourite language: ")
	fmt.Scan(&language)

	fmt.Printf("Favorite Language: % s", language)

	// Using Scan() to take multiple inputs from the user
	var name string
	var age int

	// take name and age input
	fmt.Println("Enter your name and age:")
	// fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)         // Takes values up to the new line
	fmt.Scanf("%s %d", &name, &age) // Takes input using format specifiers

	// print input values
	fmt.Printf("Name: %s\nAge: %d", name, age)

	// take float input
	var temperature float32
	fmt.Println("Enter the current temperature: ")
	fmt.Scanf("%g", &temperature)

	// take boolean input
	var sunny bool
	fmt.Println("Is the day sunny?")
	fmt.Scanf("%t", &sunny)

	fmt.Printf("Current temperature: %g\nIs the day Sunny? %t", temperature, sunny)
}
