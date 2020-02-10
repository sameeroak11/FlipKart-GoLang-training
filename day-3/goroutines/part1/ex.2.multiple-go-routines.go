package main


import (
	"fmt"
	"time"
	"runtime"
	"sync"
)


func f1(pWG *sync.WaitGroup, sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	time.Sleep(sleepMs * time.Millisecond)

	pWG.Done()

	return
}


func f2(pWG *sync.WaitGroup, sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	time.Sleep(sleepMs * time.Millisecond)

	pWG.Done()

	return
}


func f3(pWG *sync.WaitGroup, sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	time.Sleep(sleepMs * time.Millisecond)

	pWG.Done()

	return
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use
	var wg sync.WaitGroup
	//wg.Add(3)

	wg.Add(1)
	go f3(&wg, 50, "this's a message for f3().")

	wg.Add(1)
	go f2(&wg, 250, "this's a message for f2().")

	wg.Add(1)
	go f1(&wg, 100, "this's a message for f1().")

	wg.Wait()

	fmt.Printf("this's main()\n")
	//time.Sleep(9001 * time.Nanosecond)
	// time.Sleep(3 * time.Second)
	return
}
