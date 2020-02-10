package main


import (
	"fmt"
)


func main() {
	ch := make(chan string, 1)
	//ch := make(chan string)
	ch <- "sagacity"
	//ch <- "opensoach"
	//fmt.Println(<-ch)
	fmt.Println(<-ch)
}
