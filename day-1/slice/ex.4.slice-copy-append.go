package main

import (
	"fmt"
)


func main() {
	s1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	var s2 []int
	s3 := make([]int, 5, 9)
	s4 := []int{100, 110, 120, 130}

	//s2[0] = 10000


	// before copying 
	fmt.Printf("s1: %#v,  len(s1): %d,  cap(s1): %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %#v,  len(s2): %d,  cap(s2): %d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: %#v,  len(s3): %d,  cap(s3): %d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4: %#v,  len(s4): %d,  cap(s4): %d\n\n", s4, len(s4), cap(s4))

	// copying the slices 
	fmt.Printf("copying s1 to s2:\n")
	cnt1 := copy(s2, s1)
	fmt.Printf("number of elements copied: %d\n", cnt1)
	fmt.Printf("no effect on s2: %#v,  len(s2): %d,  cap(s2): %d\n\n", s2, len(s2), cap(s2))

	fmt.Printf("appending individual values to s2:\n")
	s2 = append(s2, 140, 150)
	fmt.Printf("s2.append.2: %#v,  len(s2): %d,  cap(s2): %d\n\n", s2, len(s2), cap(s2))
	//fmt.Printf("note: appending individual elements to a target slice, however, doesn't increase capacity of the target slice.\n\n")


	fmt.Printf("copying s1 to s3.\n") // appending a slice to a slice increases capacity of the target slice:\n")
	cnt2 := copy(s3, s1)
	fmt.Printf("number of elements copied: %d\n", cnt2)
	fmt.Printf("s3: %#v,  len(s3): %d,  cap(s3): %d\n\n", s3, len(s3), cap(s3))


	fmt.Printf("appending s1 to s3:\n")
	s3 = append(s3, s1...)
	fmt.Printf("s3.append.1: %#v,  len(s3): %d,  cap(s3): %d\n", s3, len(s3), cap(s3))
	//fmt.Printf("note: appending a source slice to a target slice increases capacity of the target slice.\n\n")


	fmt.Printf("wo lhs: appending individual values to s3:\n")
	s3 = append(s3, 140, 150, 1000, 2000, 3000)
	fmt.Printf("s3.append.2: %#v,  len(s3): %d,  cap(s3): %d\n", s3, len(s3), cap(s3))
	//fmt.Printf("note: appending individual elements to a target slice, however, doesn't increase capacity of the target slice.\n\n")


	/*
	fmt.Printf("copying s1 to s4:\n")
	cnt3 := copy(s4, s1)
	fmt.Println("number of elements copied:", cnt3)
	fmt.Printf("s4: %#v,  len(s4): %d,  cap(s4): %d\n\n", s4, len(s4), cap(s4))


	// s4 has been copied and hence modified permanently, i.e., s4 contains [58 69 40 45]
	fmt.Printf("copying s4 to s1:\n")
	cnt4:= copy(s1, s4)
	fmt.Println("number of elements copied:", cnt4)
	fmt.Printf("s1: %#v,  len(s1): %d,  cap(s1): %d\n\n", s1, len(s1), cap(s1))


	// appending s1 to s2
	fmt.Printf("appending s1 to s2:\n")
	s2 = append(s2, s1...)
	fmt.Printf("s2.append.1: %#v,  len(s2): %d,  cap(s2): %d\n", s2, len(s2), cap(s2))
	fmt.Printf("note: appending a source slice to a target slice increases capacity of the target slice.\n\n")
	*/
}
