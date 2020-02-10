package main


import (
	"fmt"
)

type EmptyInterface interface{}


//func printDetails(iObj interface{}) {
func printDetails(iObj EmptyInterface) {
	var intVal int
	var str string

	fmt.Printf("interface: type(iObj): %T,  value(iObj): %v\n", iObj, iObj)

	str = iObj.(string)
	fmt.Printf("concreate-type: type(intVal): %T,  value(intVal): %v\n", str, str)

	/*
	intVal, isOK := iObj.(int)
	if !isOK {
		fmt.Printf("caught type-assertion-panic, value: %v.\n", intVal)
		return
	}
	*/
	intVal = iObj.(int)
	fmt.Printf("concreate-type: type(intVal): %T,  value(intVal): %v\n", intVal, intVal)

	return
}


func main() {
	/*
	intVal := 10
	printDetails(intVal)
	*/

	strVal := "this's test."
	printDetails(strVal)

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
