package main

import "fmt"

func main() {
	f1("o/p 1")
	f1("o/p 2", "this", "is")
	f1("o/p 3", "this", "is", "a")
	f1("o/p 4", "this", "is", "a", "variadic")
	f1("o/p 5", "this", "is", "a", "variadic", "function")
	f1("o/p 6", "this", "is", "a", "variadic", "function", "example")
}

func f1(msg string, str ...string) {
	fmt.Println(msg, ":", str)
}
