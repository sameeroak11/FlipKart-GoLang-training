package main


import (
	"fmt"
	"time"
)


func f1(ch chan string) {
	for i :=0; ;i++ {
		ch <- "value-1"
		fmt.Printf("[enq]: from i-iteration[%d]: cap: %d, len: %d, \n", i, cap(ch), len(ch))
		time.Sleep(time.Second * 2)
	}
}


func f2(ch chan string) {
	for j :=0; ;j++ {
		ch <- "value-2"
		fmt.Printf("[enq]: from j-iteration[%d]: cap: %d, len: %d, \n", j, cap(ch), len(ch))
		time.Sleep(time.Second * 3)
	}
}


func f3(ch chan string) {
	for k := 0; ;k++ {
		select {
			case msg1 := <-ch:
				fmt.Printf("[deq]: from k-iteration[%d]: cap: %d, len: %d, \n", k, cap(ch), len(ch))
				fmt.Printf("msg1: %s\n", msg1)
				break

		default:	// default case is executed immediately if there're no data at any of the channels
			time.Sleep(time.Second * 1)
			fmt.Printf("no data available.\n")
			break
		}
	}
}


func main() {
	ch := make(chan string, 2)

	go f1(ch)
	go f2(ch)
	go f3(ch)

	time.Sleep(10 * time.Second)
}
