// updates returning named parameter value.


package main

import (
	"fmt"
)


func f1() (counter int) {
	counter = 10

	defer func() {
		counter++
	}()

	return 1
}


func main() {
	fmt.Printf("f1(): %d\n", f1())
}
