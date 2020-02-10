package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	/*fmt.Println(<-ch)	// naveen is out, 1 space is created.
	ch <- "steve"	// steve is in
	fmt.Println(<-ch)	// paul is out
	fmt.Println(<-ch)	// steve is out
	*/

	for val := range ch {
		fmt.Printf("val: %s\n", val)
	}
}
