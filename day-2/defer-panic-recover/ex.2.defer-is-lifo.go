// defer works in lifo order.


package main

import (
	"fmt"
)


func f1() {
	n := 10
	defer fmt.Printf("1st n: %d\n", n)
	n++
	defer fmt.Printf("2nd n: %d\n", n)
	return
}


func main() {
	f1()
}
