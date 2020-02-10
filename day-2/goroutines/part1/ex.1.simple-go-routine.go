package main


import (
	"fmt"
	"time"
)


func f1(msg string) {
	fmt.Printf("%s\n", msg)
	return
}


func main() {
	go f1("this's a message.")
	//f1("this's a message.")
	fmt.Printf("this's main()\n")
	time.Sleep(1 * time.Second)
	return
}
