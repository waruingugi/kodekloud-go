package main

import (
	"fmt"
	"sort"
)

func main() {
	// vars := []int{5, 2, 0, 3, 4, 9, 6}
	vars := []string{"Learning", "Golang", "on", "kodekloud"}
	sort.Strings(vars)
	fmt.Println(vars)
}
