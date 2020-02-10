package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)


/*func f1(done chan bool) {
	done <- true
	time.Sleep(3 * time.Millisecond)
	fmt.Printf("f1() goroutine\n")
}*/

func f1(pWG *sync.WaitGroup, done chan bool) {
	time.Sleep(2 * time.Second)
	val := <-done
	fmt.Printf("f1() goroutine:  val: %t\n", val)

	done <- false

	pWG.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use

	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan bool)
	go f1(&wg, done)
	done <- true

	fmt.Println("main() goroutine:", <-done)
	wg.Wait()
	//<-done
}
