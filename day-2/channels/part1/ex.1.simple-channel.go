package main

import (
	"fmt"
	"time"
	"runtime"
	//"sync"
)


/*func f1(done chan bool) {
	done <- true
	time.Sleep(3 * time.Millisecond)
	fmt.Printf("f1() goroutine\n")
}*/

func f1(done chan bool) {
	val := <-done
	fmt.Printf("f1() goroutine:  val: %t\n", val)
	done <- false
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use

	// var wg sync.WaitGroup
	// wg.Add(2)

	done := make(chan bool)
	go f1(done)
	done <- true
	time.Sleep(1 * time.Second)
	fmt.Println("main() goroutine:", <-done)
	//<-done
}
