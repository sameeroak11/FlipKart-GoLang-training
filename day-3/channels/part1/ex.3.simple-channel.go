package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("f1():  Hello world goroutine")
	fmt.Printf("f1():  flag: %t\n", <-done)
	//done <- true
	fmt.Println("f1():  test.")
}

func main() {
	done := make(chan bool)

	go hello(done)
	done <- true
	//fmt.Println("main():  done got back:", <-done)
	time.Sleep(2 * time.Second)
	fmt.Println("main():  main function")
}
