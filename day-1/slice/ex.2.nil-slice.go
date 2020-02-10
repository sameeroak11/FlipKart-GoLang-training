package main

import (
	"fmt"
)


func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil slice.")
	}

	s = []int {100, 200}
	if s != nil {
		fmt.Println("slice is not nil.")
		fmt.Printf("len(s): %d,  cap(s): %d\n", len(s), cap(s))
	}
}
