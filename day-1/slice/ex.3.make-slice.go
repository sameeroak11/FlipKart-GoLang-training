package main

import (
	"fmt"
)


func main() {
	s1 := []int {10, 20, 30, 40, 50}
	//s1 := make([]int, 5)
	printSlice("1: s1->", s1)

	//s2 := make([]int, 0, 5)
	s2 := []int {100, 200, 300, 400, 500}
	printSlice("2: s2->", s2)

	s3 := s2[:2]
	printSlice("3: s3->", s3)

	s4 := s3[0:4]
	printSlice("4: s4->", s4)

	/* s1 = s1[:7]
	printSlice("5: panic s1->", s1) */

	s5 := s3[0:]
	printSlice("6: s5->", s5)

	s6 := s3[:len(s3)]
	printSlice("7: s6->", s6)

	s7 := s1[:]
	printSlice("8: s7->", s7)
}

func printSlice(msg string, s []int) {
	fmt.Printf("%s    len: %d,   cap: %d,    slice: %#v\n", msg, len(s), cap(s), s)
}
