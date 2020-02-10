package main

import "fmt"


func f1() {
	/* defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in f():", r)
		}
	}()*/
	fmt.Println("invoking f2.")
	f2(0)
	fmt.Println("f2() default return.")
}


func f2(i int) {
	if i > 3 {
		fmt.Println("panic in f2():", i)
		panic(fmt.Sprintf("%d", i))
	}
	defer fmt.Println("defer f2():", i)
	fmt.Println("i in f2():", i)
	f2(i + 1)
}


func main() {
	f1()
	fmt.Println("default return from f1().")
}
