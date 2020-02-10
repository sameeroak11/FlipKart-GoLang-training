package main


import (
	"fmt"
	"time"
)


func f1(sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	time.Sleep(sleepMs * time.Millisecond)
	return
}


func f2(sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	time.Sleep(sleepMs * time.Millisecond)
	return
}


func f3(sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	time.Sleep(sleepMs * time.Millisecond)
	return
}


func main() {
	go f3(50, "this's a message for f3().")
	go f2(250, "this's a message for f2().")
	go f1(100, "this's a message for f1().")
	fmt.Printf("this's main()\n")
	//time.Sleep(9001 * time.Nanosecond)
	time.Sleep(3 * time.Second)
	return
}
