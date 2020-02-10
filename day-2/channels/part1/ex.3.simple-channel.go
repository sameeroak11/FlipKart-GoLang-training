package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	<-done //<- true
	//done <- true
}
func main() {
	done := make(chan bool)

	go hello(done)
	done <- true
	fmt.Println("done got back:", <-done)
	time.Sleep(2 * time.Second)
	fmt.Println("main function")
}
