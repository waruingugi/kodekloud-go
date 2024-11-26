// A struct definition is just a blueprint.
// To use a struct we need to to create an instance of it

package main

import "fmt"

// initalize the function Rectangle
type Rectangle func(int, int) int

// create structure
type rectanglePara struct {
	length  int
	breadth int
	color   string

	// function as a field of struct
	rect Rectangle
}

func main() {
	// // declare a struct
	// type Person struct {
	// 	name string
	// 	age  int
	// }

	// // assign value to struct while creating an instance
	// person1 := Person{"John", 25}
	// fmt.Println(person1)

	// // define an instance
	// var person2 Person

	// // assign value to struct variables
	// person2 = Person{
	// 	name: "Sar",
	// 	age:  29,
	// }
	// fmt.Println(person2)

	// // declare a struct
	// type Rectangle struct {
	// 	length  int
	// 	breadth int
	// }

	// // declare instance rect1 and defining the struct
	// rect := Rectangle{22, 12}

	// // access the bread of the struct
	// fmt.Println("Breadth:", rect.breadth)

	// area := rect.length * rect.breadth
	// fmt.Println("Area: ", area)

	// assign values to struct variables
	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",

		rect: func(length int, breadth int) int {
			return length * breadth
		},
	}

	fmt.Println("Color of Rectangle: ", result.color)
	fmt.Println("Area of Rectangle: ", result.rect(result.length, result.breadth))

}
