/*
Type casting is the process of converting the value
of one data type to another data type.
*/
package main

import "fmt"

func main() {
	var floatValue float32 = 5.45

	// type conversion from float to int
	var intValue int = int(floatValue)

	fmt.Printf("Float Value is %g\n", floatValue)
	fmt.Printf("Integer Value is %d", intValue)

	// In implicit type casting , one data type is automatically
	// converted to another data type. But Go doesn't support implicit type casting.
	// var number int = 4.34  // This will raise an error

	var number1 int = 20
	var number2 float32 = 5.7
	var sum float32

	// addition of different data types
	sum = float32(number1) + number2
	fmt.Printf("Sum is %g", sum)

}
