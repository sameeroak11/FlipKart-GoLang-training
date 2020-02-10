package main


import (
	"fmt"
)


func f1(ch1 chan int) {
	val := <-ch1
	fmt.Printf("channel value: %d\n", val)
	return
}


func main() {
	//ch := make(chan int, 1)
	ch := make(chan int)
	go f1(ch)
	ch <- 5
	// <-ch

	fmt.Printf("main proceeds.\n")
}
