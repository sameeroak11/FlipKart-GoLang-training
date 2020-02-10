package main


import ( "fmt"
	"time"
	"runtime"
	"sync"
)


func f1(pWg *sync.WaitGroup, sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	time.Sleep(sleepMs * time.Millisecond)
	pWg.Done()
	return
}


func f2(pWg *sync.WaitGroup, sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	time.Sleep(sleepMs * time.Millisecond)
	pWg.Done()
	return
}


func f3(pWg *sync.WaitGroup, sleepMs time.Duration, msg string) {
	fmt.Printf("[%v]: %s\n", sleepMs, msg)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("goroutine f3() in iteration: %d\n", i)
	}
	pWg.Done()
	return
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use
	var wg sync.WaitGroup

	wg.Add(1)
	go func(pWg *sync.WaitGroup) {
		fmt.Printf("this's in closure. and am sleeping for 5 seconds.\n")
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("goroutine closure in iteration: %d\n", i)
		}

		pWg.Done()
		return
	}(&wg)

	wg.Add(1)
	go f3(&wg, 50, "this's a message for f3().")

	wg.Add(1)
	go f2(&wg, 250, "this's a message for f2().")

	wg.Add(1)
	go f1(&wg, 100, "this's a message for f1().")

	wg.Wait()		// resembles join in other programming languages, for instance, java, c (using pthreads)
	fmt.Printf("this's main()\n")
	return
}
