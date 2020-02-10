package main


import (
	"fmt"
)


func main() {
	ch := make(chan string, 1)
	//ch := make(chan string)
	ch <- "sagacity"
	<-ch
	ch <- "opensoach"
	fmt.Println(<-ch)
}
