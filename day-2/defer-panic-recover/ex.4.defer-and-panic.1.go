package main

import (
	"fmt"
)


func f1() {
	defer func() {
		if recoveredVal := recover(); recoveredVal != nil {
			fmt.Println("f1():  recover value:", recoveredVal)
		} else {
			fmt.Println("f1():  no panic")
		}
	}()

	fmt.Println("f1():  invoking f2.")
	f2()
	fmt.Println("f1():  f2() returns.")
}


func f2() {
	intArr := []int{10, 20, 30, 40}
	fmt.Printf("f2():  intArr[3]: %d\n", intArr[3])
	//fmt.Printf("f2():  intArr[4]: %d\n", intArr[4])
	panic(intArr[0])
	panic(intArr[1])
}


func main() {
	f1()
	fmt.Printf("main():  default return.\n")
}
