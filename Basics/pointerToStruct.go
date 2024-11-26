package main

import "fmt"

func main() {
	// // declare a struct Person
	// type Person struct {
	// 	name string
	// 	age  int
	// }

	// // instance of the struct Person
	// person1 := Person{"John", 25}

	// // create a struct type pointer that
	// // stores the address of person 1
	// var ptr *Person
	// ptr = &person1

	// fmt.Println(person1)
	// fmt.Println(ptr)

	// // access the name member
	// fmt.Println("Name: ", ptr.name)
	// fmt.Println("Age:", ptr.age)

	// create struct variable
	type Weather struct {
		city        string
		temperature int
	}
	weather := Weather{"California", 20}
	fmt.Println("Initial Weather:", weather)

	// create struct type pointer
	ptr := &weather

	// change value of temperature to 25
	ptr.temperature = 25

	fmt.Println("Updated Weather:", weather)
}
