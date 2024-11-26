package main

import "fmt"

func main() {
	// numbers := [5]int{21, 24, 27, 30, 33}

	// for index, item := range numbers {
	// 	fmt.Printf("numbers[%d] = %d \n", index, item)
	// }

	// string := "Golang"
	// fmt.Println("Index: Character")

	// // i access index of each character
	// // item access each character
	// for i, item := range string {
	// 	fmt.Printf("%d= %c \n", i, item)
	// }

	// // create a map
	// subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	// fmt.Println("Marks obtained: ")

	// for subject, marks := range subjectMarks {
	// 	fmt.Println(subject, ":", marks)
	// }

	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}

	fmt.Println("Subjects:")
	for subject := range subjectMarks {
		fmt.Println(subject)
	}
}
