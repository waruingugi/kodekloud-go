package main

import "fmt"

func main() {
	// dayOfWeek := 3

	// switch dayOfWeek {

	// case 1:
	// 	fmt.Println("Sunday")

	// case 2:
	// 	fmt.Println("Monday")

	// case 3:
	// 	fmt.Println("Tuesday")
	// 	fallthrough // Use fallthrough if we need to execute other cases after the matching case

	// case 4:
	// 	fmt.Println("Wednesday")

	// case 5:
	// 	fmt.Println("Thursday")

	// case 6:
	// 	fmt.Println("Friday")

	// default:
	// 	fmt.Println("Invalid day")
	// }

	// dayOfWeek := "Sunday"

	// switch dayOfWeek {
	// case "Saturday", "Sunday":
	// 	fmt.Println("Weekend")

	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("Weekday")

	// default:
	// 	fmt.Println("Invalid day")
	// }

	switch day := 4; day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid Day!")
	}
}
