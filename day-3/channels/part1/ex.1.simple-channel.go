package main


import (
	"fmt"
	"time"
)


func f1(ch1 chan int) {
	val := <-ch1
	fmt.Printf("f1():  channel value: %d\n", val)
	return
}


func main() {
	ch := make(chan int)
	go f1(ch)
	time.Sleep(2 * time.Second)
	ch <- 5

	fmt.Printf("main():  main proceeds.\n")
}
