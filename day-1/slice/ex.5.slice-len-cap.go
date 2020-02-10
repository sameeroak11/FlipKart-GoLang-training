package main

import (
	"fmt"
)


func main() {
	fmt.Println("1: original slice.")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// slice sliced to zero length.
	fmt.Println("2: length truncated to 0.")
	s = s[:0]
	printSlice(s)

	// length extended.
	fmt.Println("3: now extended to 4.")
	s = s[:4]
	printSlice(s)

	// dropped first 2 elements.
	fmt.Println("4: first 2 values dropped.")
	s = s[2:]
	printSlice(s)

	// 1 more element in slice. should panic.
	fmt.Println("5: 1 more element in s. should panic.")
	s[2] = 1000
	printSlice(s)

	// attempting to extend to 5. should panic now, too.
	fmt.Println("6: extended to 5, should panic now, too.")
	s = s[:5]
	printSlice(s)
}


func printSlice(s []int) {
	fmt.Printf("len: %d,  cap: %d,  slice: %v\n\n", len(s), cap(s), s)
}
