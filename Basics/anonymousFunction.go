package main

import "fmt"

// regular function to calculate square of numbers
func findSquare(num int) int {
	square := num * num
	return square
}

func displayNumber() func() int {
	number := 10
	return func() int {
		number++
		return number
	}
}

func main() {
	// // anonymous function
	// var greet = func() {
	// 	fmt.Println("Hello, how are you")
	// }

	// // function call
	// greet()

	// // anonymous function with arguments
	// var sum = func(n1, n2 int) {
	// 	sum := n1 + n2
	// 	fmt.Println("Sum is:", sum)

	// 	// return sum
	// }

	// sum(5, 3)

	// area := func(length, breadth int) int {
	// 	return length * breadth
	// }

	// fmt.Println("The area of rectangle is", area(3, 4))

	// // anonymous function that returns sum of numbers
	// sum := func(number1 int, number2 int) int {
	// 	return number1 + number2
	// }

	// // function call
	// result := findSquare(sum(6, 9))
	// fmt.Println("Result is:", result)

	a := displayNumber()
	fmt.Println(a())
}
