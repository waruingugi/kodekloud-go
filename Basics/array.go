package main

import "fmt"

func main() {
	// declare array variable of type integer
	// defined size [size=5]
	var arrayOfInteger = [5]int{1, 5, 6, 0, 3}

	fmt.Println(arrayOfInteger)

	// Declaring an array without specifying it's size
	var arrayOfString = [...]string{"Hello", "Programiz"}
	fmt.Println(arrayOfString)

	languages := [3]string{"Go", "Java", "C++"}

	// access element at index 0
	fmt.Println(languages[0])

	// // decalre an array
	// var arrayOfIntegers [3]int

	// elements are assigned using index
	// arrayOfIntegers[0] = 5
	// arrayOfIntegers[1] = 10
	// arrayOfIntegers[2] = 15

	// fmt.Println(arrayOfIntegers)

	// initialize the elements of index 0 and 3 only
	// arrayOfIntegers := [5]int{0: 7, 3: 9}
	// fmt.Println(arrayOfIntegers)

	// weather := [3]string{"Rainy", "Sunny", "Cloudy"}

	// // Change the element at index 2
	// weather[2] = "Stromy"
	// fmt.Println(weather)

	// create an array
	// var arrayOfIntegers = [...]int{1, 5, 8, 0, 3, 10}

	// // find the length of array using len()
	// length := len(arrayOfIntegers)

	// fmt.Println("The length of array is", length)

	// age := [...]int{12, 4, 5}

	// // loop through the array
	// for i := 0; i < len(age); i++ {
	// 	fmt.Println(age[i])
	// }

	// create a 2 dimensional array
	arrayInteger := [2][2]int{{1, 2}, {3, 4}}

	// access the values of 2d array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arrayInteger[i][j])
		}
	}
}
