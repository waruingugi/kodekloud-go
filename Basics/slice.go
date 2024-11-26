// A slice is a collection of similar types of data, just like an array
// But, unlike arrays, slice doesn't have a fixed size.
// Arrays have a fixed size but slices do not.
// Slices also have a [] notation because we can add or remove elements therefore changing it's size

package main

import "fmt"

func main() {
	// // declare slice variable of type integer
	// numbers := []int{1, 2, 3, 4, 5}

	// // print the slice
	// fmt.Println("Numbers: ", numbers)

	// // an integer array
	// numbers := [8]int{10, 20, 30, 40, 50, 60, 70, 80}

	// // create a slice from an array
	// sliceNumbers := numbers[4:7]
	// fmt.Println(sliceNumbers)

	// // Adding elements to a slice
	// primeNumbers := []int{2, 3}

	// // add elements 5, 7 to the slice
	// primeNumbers = append(primeNumbers, 5, 7)
	// fmt.Println("Prime Numbers:", primeNumbers)

	// create two slice
	// evenNumbers := []int{2, 4}
	// oddNumbers := []int{1, 3}

	// evenNumbers = append(evenNumbers, oddNumbers...)
	// fmt.Println("Numbers: ", evenNumbers)

	// create two slices
	// primeNumbers := []int{2, 3, 5, 7}
	// numbers := []int{1, 2, 3}

	// // copy elements of primeNumbers to numbers
	// copy(numbers, primeNumbers)

	// // print numbers
	// fmt.Println("Numbers:", numbers)

	// // find the length of the slice
	// length := len(numbers)

	// fmt.Println("Length:", length)

	// numbers := []int{2, 4, 6, 8, 10}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// create a slice using make() function
	numbers := make([]int, 5, 7)
	fmt.Println("Numbers: ", numbers)
}
