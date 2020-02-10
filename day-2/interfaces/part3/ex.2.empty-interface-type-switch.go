package main


import (
	"fmt"
)

type EmptyInterface interface{}


//func printDetails(iObj interface{}) {
func printDetails(iObj EmptyInterface) {
	switch iObj.(type) {
		case int:
			intVal := iObj.(int)
			fmt.Printf("[int]: interface: type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
			fmt.Printf("type: %T,  value(intVal): %v\n", intVal, intVal)
			break
		case string:
			strVal := iObj.(string)
			fmt.Printf("[string]: interface: type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
			fmt.Printf("type: %T,  value(strVal): %v\n", strVal, strVal)
			break
		case MyTypeObj:
			myTypeObj := iObj.(MyTypeObj)
			fmt.Printf("[MyTypeObj]: interface: type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
			fmt.Printf("type: %T,  value(myTypeObj): %v\n", myTypeObj, myTypeObj)
			break
		case *int:
			pIntVal := iObj.(*int)
			fmt.Printf("[int]: interface: type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
			fmt.Printf("type: %T,  value(pIntVal): %v\n", pIntVal, *pIntVal)

		default:
			fmt.Printf("[default]: interface: type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
			break
	}

	return
}

type MyTypeObj struct {
	Name string
}

func main() {
	strVal := "this's test."
	printDetails(strVal)

	intVal := 10
	printDetails(intVal)

	intVal1 := 20
	printDetails(&intVal1)

	myTypeObj := MyTypeObj {
		Name: "infinity",
	}
	printDetails(myTypeObj)

	printDetails(10.10)

	return
}
