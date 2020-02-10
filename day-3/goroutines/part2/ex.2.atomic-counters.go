package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func main() {
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("counter: %d", counter)
}
