// length-and-capacity.go


package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)

	ch <- "sagacity"
	ch <- "opensoach"

	fmt.Printf("capacity: %d,  length: %d\n", cap(ch), len(ch))
	fmt.Printf("read value-1: %s\n", <-ch)
	fmt.Printf("new capacity: %d,  new length: %d\n", cap(ch), len(ch))
	fmt.Printf("read value-2: %s\n", <-ch)
	fmt.Printf("new capacity: %d,  new length: %d\n", cap(ch), len(ch))
	fmt.Printf("read value-3: %s\n", <-ch)
}
