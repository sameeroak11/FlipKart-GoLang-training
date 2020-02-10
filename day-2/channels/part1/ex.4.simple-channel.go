package main


import (
	"fmt"
	//"time"
)


func f1(done chan bool) {
	done <- true
	//time.Sleep(5 * time.Second)
	fmt.Printf("this's f1() goroutine\n")
}

func f2(done chan bool) {
	val := <-done
	fmt.Printf("this's f2() goroutine. val: %t\n", val)
}


func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go f1(done1)
	// <-done
	fmt.Printf("main goroutine: %t\n", <-done1)

	go f2(done2)
	done2 <- true

	/* var chan1 chan int
	if chan1 == nil {
		fmt.Printf("chan1 is %v\n", chan1)
	} else {
		fmt.Printf("chan1 is non-nil: %v\n", chan1) */
}
