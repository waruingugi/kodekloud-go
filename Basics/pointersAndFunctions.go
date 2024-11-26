package main

import "fmt"

// // function defintion with a pointer argument
// func update(num *int) {
// 	*num = 30
// }

// func main() {
// 	var number = 55

// 	update(&number)

// 	fmt.Println("The number is", number)
// }

// func main() {
// 	result := display()
// 	fmt.Println("Welcome to", *result)
// }

// func display() *string {
// 	message := "Programiz"

// 	return &message
// }

// call by value
func callByValue(num int) {
	num = 30
	fmt.Println(num) // 30
}

// cally by reference
func callyByReference(num *int) {
	*num = 10
	fmt.Println(*num) // 10
}

func main() {
	var number int

	// passing value
	callByValue(number)

	// passing a reference (address)
	callyByReference((&number))
}
