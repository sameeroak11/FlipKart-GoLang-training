// centralize control of decentralized services


package main

import (
	"fmt"
	"runtime"
	"time"
	//"bufio"
	//"os"
	//"strings"
	"math/rand"
	"sync"
)


func t1(ch1 <-chan int, ch chan<-int, ich <-chan int) {
	var i int
	for i = 0;;i++ {
		select {
			case val := <-ch1:
				fmt.Printf("[ich: %d][%d: thread-t1]: val: %d\n", <-ich, i, val)
				ch <- 1
				break
		}
	}

	fmt.Printf("[%d: thread-t1]\n", i)
}

func t2(ch1 <-chan int, ch chan<-int, ich <-chan int) {
	var i int
	for i = 0;;i++ {
		select {
			case val := <-ch1:
				fmt.Printf("[ich: %d][%d: thread-t2]: val: %d\n", <-ich, i, val)
				ch <- 2
				break
		}
	}

	fmt.Printf("[%d: thread-t2]\n", i)
}

func t3(ch1 <-chan int, ch chan<-int, ich <-chan int) {
	var i int
	for i = 0;;i++ {
		select {
			case val := <-ch1:
				fmt.Printf("[ich: %d][%d: thread-t3]: val: %d\n", <-ich, i, val)
				ch <- 3
				break
		}
	}

	fmt.Printf("[%d: thread-t3]\n", i)
}


func randomInt(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ich := make(chan int)

	go t1(ch1, ch, ich)
	go t2(ch2, ch, ich)
	go t3(ch3, ch, ich)

	rand.Seed(time.Now().UnixNano())

	go func(pWG *sync.WaitGroup) {
		//for i := 0; i < 20; i++ {
		for i := 0; i < 4; i++ {
			randInt := randomInt(0, 1000)
			time.Sleep(time.Duration(randInt) * time.Millisecond)
			//fmt.Printf("[i: %d, %d]\n", i, randInt)

			select {
				case signal := <-ch:
					switch signal {
						case 0:
							//fmt.Printf("[i: %d, %d] signal: %d\n", i, randInt, signal)
							ch1 <- 1
							ich <- i
							break
						case 1:
							//fmt.Printf("[i: %d, %d] signal: %d\n", i, randInt, signal)
							ch2 <- 2
							ich <- i
							break
						case 2:
							//fmt.Printf("[i: %d, %d] signal: %d\n", i, randInt, signal)
							ch3 <- 3
							ich <- i
							break
						case 3:
							//fmt.Printf("[i: %d, %d] signal: %d\n", i, randInt, signal)
							go func() {
								ch <- 0
							}()
							break
					}
					break
			}
		}

		pWG.Done()
	}(&wg)
	ch <- 0

	wg.Wait()

	/*
	time.Sleep(time.Second * 2)
	var strInput string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nUser response:\n")
	for {
		fmt.Printf("Do you want to terminate the application: ")
		strInput, _ = reader.ReadString('\n')

		// converts CRLF to LF
		strInput = strings.Replace(strInput, "\n", "", -1)

		switch strInput {
			case "y", "Y", "yes", "Yes", "YES":
				fmt.Printf("Server terminating on user response: \"%s\"\n", strInput)
				time.Sleep(time.Second * 3)  // waits for 3 seconds so that logger can push pending messages, though not guaranteed :).
				os.Exit(0)
		}
	}
	*/
}
