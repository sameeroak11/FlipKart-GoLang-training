package main


import (
	"fmt"
	"time"
)


func f1(secondsSleep time.Duration, done chan bool) {
	fmt.Printf("this's f1() goroutine\n")
	done <- true
	time.Sleep(secondsSleep * time.Second)
	fmt.Printf("this's f1() goroutine, after %d seconds sleep\n", secondsSleep)
	return
}


func main() {
	done := make(chan bool)
	go f1(1, done)
	<-done
	fmt.Printf("main goroutine\n")
	//secondsSleep := 10
	//time.Sleep(time.Duration(secondsSleep) * time.Second)
	//fmt.Printf("main goroutine after %d seconds sleep\n", secondsSleep)
}
