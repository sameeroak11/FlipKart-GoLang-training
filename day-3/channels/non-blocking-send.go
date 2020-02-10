// non-blocking send
package main

import (
	"fmt"
)


func main() {
	var ok bool
	...
	select {
		case ch <- item:
			ok = true
		default:
			ok = false
	}

	// at this point, ok is:
	//   true  => enqueued without blocking
	//   false => not enqueued, would've blocked since channel was full.
}
