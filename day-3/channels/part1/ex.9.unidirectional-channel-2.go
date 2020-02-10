// it's possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.

package main

import (
	"fmt"
)

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
}
