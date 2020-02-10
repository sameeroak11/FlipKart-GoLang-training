package main

/* i := 0
for ; i < 10; i++ {
	...
	...
}
fmt.Println("i:", i)


i := 0
...

for ; i < 10; i++ {

}


for ;; i++ {

}

for {

} */


import (
	"fmt"
)


func main() {
	n := 3

	for i := n; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}


