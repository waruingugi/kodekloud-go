package main

import "fmt"

func main() {
	// for loop terminates when i becomes 6
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// var n, sum = 10, 0
	// for i := 1; i <= n; i++ {
	// 	sum += 1
	// }

	// fmt.Println("sum =", sum)

	numbers := [5]int{11, 22, 33, 44, 55}

	for _, item := range numbers {
		fmt.Println(item)
	}

	// for item := range numbers {
	// 	fmt.Println(numbers[item])
	// }
}
