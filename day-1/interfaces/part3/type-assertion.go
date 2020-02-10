package main


import (
	"fmt"
)


func printDetails(iObj interface{}) {
	var intVal int

	intVal = iObj.(int)
	fmt.Printf("interface: type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
	fmt.Printf("concreate-type: type(intVal): %T,  value(intVal): %v\n", intVal, intVal)
}


func main() {
	intVal := 10
	printDetails(intVal)

	//strVal := "this's test."
	//printDetails(strVal)

	/*
	myTypeObj := struct {
		Name string
	} {
		Name: "infinity",
	}

	printDetails(myTypeObj)
	*/

	return
}
