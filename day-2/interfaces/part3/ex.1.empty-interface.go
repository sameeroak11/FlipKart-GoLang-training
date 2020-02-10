package main


import (
	"fmt"
)


func printDetails(iObj interface{}) {
	fmt.Printf("type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
}


func main() {
	intVal := 10
	printDetails(intVal)

	strVal := "this's test."
	printDetails(strVal)

	myTypeObj := struct {
		Name string
	} {
		Name: "infinity",
	}

	printDetails(myTypeObj)

	return
}
