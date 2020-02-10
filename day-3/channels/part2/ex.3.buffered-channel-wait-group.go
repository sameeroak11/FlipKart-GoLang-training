package main


import (
	"fmt"
	"sync"
	"time"
)


/*
it's mandatory to pass the address of wg. unless, each goroutine has its own copy of the WaitGroup and main goroutine won't be
notified when all of them finish execution. */
func process(i int, pWg *sync.WaitGroup) {
	fmt.Printf("started goroutine: %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine %d done\n", i)
	pWg.Done()
}


func main() {
	no := 3
	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()

	fmt.Printf("all go-routines are done on their respective part.\n")
}
