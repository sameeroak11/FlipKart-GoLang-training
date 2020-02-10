package main

import (
	"fmt"
)


func main() {
	s1 := []int {10, 20, 30}
	s2 := []int{}

	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s2: %#v\n", s2)

	fmt.Printf("len(s1): %d\n", len(s1))
	fmt.Printf("len(s2): %d\n", len(s2))

	fmt.Printf("len(s1): %d,  cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d,  cap(s2): %d\n", len(s2), cap(s2))
}
