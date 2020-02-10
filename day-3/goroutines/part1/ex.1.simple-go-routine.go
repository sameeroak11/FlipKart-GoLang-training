package main


import (
	"fmt"
	"time"
	"runtime"
	"sync"
)


func f1(pWG *sync.WaitGroup, msg string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("f1():  %s\n", msg)

	pWG.Done()

	return
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use
	var wg sync.WaitGroup
	wg.Add(1)

	go f1(&wg, "this's a message.")
	// return

	fmt.Printf("main():  this's main()\n")

	//wg.Wait()
	return
}
