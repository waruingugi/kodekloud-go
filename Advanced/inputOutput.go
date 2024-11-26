// package main

// import "fmt"

// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Rectangle struct {
// 	length, width float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.length * r.width
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.length + r.width)
// }

// func main() {
// 	var s Shape = Rectangle{4.0, 6.0}
// 	fmt.Println(s.Area())
// 	fmt.Println(s.Perimeter())
// }

package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Learning is fun")
	buf := make([]byte, 4)

	for {
		n, err := r.Read(buf)
		fmt.Println(buf[:n], err)

		if err != nil {
			fmt.Println("breaking out")
			break
		}

	}

}
