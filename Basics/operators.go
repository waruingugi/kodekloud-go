package main

import "fmt"

func main() {
	num1 := 6
	num2 := 2

	// + adds two variables
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num1, sum)

	// - substact two variables
	difference := num1 - num2
	fmt.Printf("%d - %d = %d\n", num1, num2, difference)

	// * multiply two variables
	product := num1 * num2
	fmt.Printf("%d * %d is %d\n", num1, num2, product)

	// / divide two integer values
	quotient := num1 / num2
	fmt.Printf("%d / %d = %d\n", num1, num2, quotient)

	// / divide two floating point variables
	result := 11.0 / 4.0
	fmt.Printf("%g / %g = %g\n", 11.0, 4.0, result)

	// % modulo-divides two variables
	remainder := num1 % num2
	fmt.Println(remainder)

	num := 5

	// increment of num by 1
	num++
	fmt.Println(num) // 6

	// decrement of num by 1
	num--
	fmt.Println(num) // 4

	// = operator to assign the value of num to result
	var resultNum int
	resultNum = num
	fmt.Println(resultNum)

	memoryNum := 20

	// &num print the memory address where 20 is stored
	fmt.Println(&memoryNum)

	// A pointer variable stores the memory variable of an address
	b := 20

	// pointer declaraation
	var numPointer *int = &b

	// gives the memory address
	fmt.Println(numPointer)
}
