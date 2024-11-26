// In Go, strings are immutable.
// This means once we create a string, we cannot change it.

package main

import (
	"fmt"
	"strings"
)

func main() {
	// //  represent string with `  `
	// message := `I love Go Programming`
	// fmt.Println(message)

	// // create and initialize a string
	// name := "Programiz"
	// fmt.Printf("%c\n", name[0])

	// // access last character
	// fmt.Printf("%c", name[8])

	// message1 := "I love"
	// message2 := "Go programming"

	// result := message1 + " " + message2
	// fmt.Println(result)

	// // create three strings
	// string1 := "Programiz"
	// string2 := "Programiz Pro"
	// string3 := "Programiz"

	// // compare strings
	// fmt.Println(strings.Compare(string1, string2)) // -1
	// fmt.Println(strings.Compare(string2, string3)) // 1
	// fmt.Println(strings.Compare(string1, string3)) // 0

	// text := "Go Programming"
	// substring1 := "Go"
	// substring2 := "Golang"

	// // check if Go is present in Go programming
	// result := strings.Contains(text, substring1)
	// fmt.Println(result)

	// // check if Golang is present in Go programming
	// result = strings.Contains(text, substring2)
	// fmt.Println(result)

	// text := "car"
	// fmt.Println("Old String:", text)

	// // replace r with t
	// replacedText := strings.Replace(text, "r", "t", 1)
	// fmt.Println("New String:", replacedText)

	text1 := "go is fun to learn"
	text1 = strings.ToUpper(text1)
	fmt.Println(text1)

	text2 := "I LOVE GOLANG"
	text2 = strings.ToLower(text2)
	fmt.Println(text2)

	var message = "I Love Golang"
	// split string from space " "
	splittedString := strings.Split(message, " ")
	fmt.Println(splittedString)

	words := []string{"I", "love", "Golang"}

	// join each element of the slice
	message = strings.Join(words, " ")
	fmt.Println(message)

	// use escape character in string
	message = "This article is about \"String\" in Go programming"
	fmt.Println(message)
}
