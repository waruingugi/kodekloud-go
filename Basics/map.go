package main

import "fmt"

func main() {
	// subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}
	// // var subjectMarks = map[string]float32{"Golang": 85, "Java": 80, "Python": 81}
	// fmt.Println(subjectMarks)

	// // create a map
	// flowerColor := map[string]string{"Sunflower": "Yellow", "Jamine": "White", "Hibiscus": "Red"}

	// // access value for key Sunflower
	// fmt.Println(flowerColor["Sunflower"]) // Yellow
	// fmt.Println(flowerColor["Hibiscus"])  // Red

	// // Create a map
	// capital := map[string]string{"Nepal": "Kathmandu", "US": "New York"}

	// capital["US"] = "Washignton DC"
	// fmt.Println("Updated Map: ", capital)

	// students := map[int]string{1: "John", 2: "Lily"}
	// students[3] = "Robin"
	// students[5] = "Julie"

	// fmt.Println("Updated Map: ", students)

	// personAge := map[string]int{
	// 	"Hermione": 21,
	// 	"Harry":    20,
	// 	"John":     25,
	// }
	// delete(personAge, "John")
	// fmt.Println("Updated Map: ", personAge)

	// squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	// // for range loop to iterate through each key value of map
	// for number, squared := range squaredNumber {
	// 	fmt.Printf("Sqaure of %d is %d\n", number, squared)
	// }

	// // Create a map using make()
	// student := make(map[int]string)

	// student[1] = "Harry"
	// student[2] = "Lilly"
	// student[5] = "Harmonie"

	// fmt.Println(student)

	// create a map
	places := map[string]string{"Nepal": "Kathmandi", "Norway": "Oslo"}

	for country := range places {
		fmt.Println(country)
	}
}
