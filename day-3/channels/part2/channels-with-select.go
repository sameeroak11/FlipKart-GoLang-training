package main


import (
	"fmt"
	"time"
)


func main() {
	c1 := make(chan string)
	//c2 := make(chan string)

	go func() {
		for {
			fmt.Printf("from c1\n")
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	/*
	go func() {
		for {
			fmt.Printf("from c2\n")
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	*/

	go func() {
		for {
			select {
				case msg1 := <-c1:
					fmt.Printf("msg1: %s\n", msg1)
					break

					/*
				case msg2 := <-c2:
					fmt.Printf("msg2: %s\n", msg2)
					break
					*/

				default:	// default case is executed immediately if there're no data at any of the channels
					time.Sleep(time.Second * 1)
					fmt.Printf("no data available.\n")
					break
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
