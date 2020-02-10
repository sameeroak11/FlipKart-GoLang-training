package main


import (
	"fmt"
	"time"
	"runtime"
	"sync"
)


func pinger(pWG *sync.WaitGroup, c chan<-string) {
	for i := 0; ; i++ {
		c <- "ping"
	}

	pWG.Done()
}


func printer(pWG *sync.WaitGroup, c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}

	pWG.Done()
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // allocates one logical processor for the scheduler to use
	var wg sync.WaitGroup
	wg.Add(2)

	var c chan string = make(chan string)
	//c := make(chan string)
	go pinger(&wg, c)
	go printer(&wg, c)

	wg.Wait()

	var input string
	fmt.Scanln(&input)
}
