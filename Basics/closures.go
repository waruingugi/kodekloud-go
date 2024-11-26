package main

import "fmt"

// outer function
func greet() func() string {
	name := "John"

	// return a nested anonymous function
	return func() string {
		name = "Hi " + name
		return name
	}
}

func calculate() func() int {
	num := 1

	// return inner function
	return func() int {
		num = num + 2
		return num
	}
}

func main() {
	// // call the outer function
	// message := greet()

	// // call the inner function
	// fmt.Println(message())

	// call the outer function
	odd := calculate()

	// call the inner function
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())

	// call the outer function again
	odd2 := calculate()
	fmt.Println(odd2())
}
