package main

import "fmt"

func main() {
	f1("this", "is", "a", "variadic", "function", "example")
}

func f1(str ...string) {
	fmt.Println(str[0])
	fmt.Println(str[3])
	fmt.Printf("type(s): %T\n", str)
}
