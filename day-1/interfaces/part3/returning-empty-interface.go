package main


import (
	"fmt"
)


type EndPoint struct {
	MAC string
	SrNo string
}

type AccessPoint struct {
	MAC string
	SrNo string
	OpMode string
}

type Creator interface {
	CreateObj() interface{}
}


func (obj EndPoint) CreateObj() interface{} {
	dummyObj := EndPoint {
		MAC: "[EndPoint.MAC]: abcd",
		SrNo: "[EndPoint.SrNo]: 012",
	}

	return interface{}(dummyObj)
}

func (obj AccessPoint) CreateObj() interface{} {
	dummyObj := AccessPoint {
		MAC: "[AccessPoint.MAC]: abcd",
		SrNo: "[AccessPoint.SrNo]: 012",
		OpMode: "[AccessPoint.OpMode]: Repeater",
	}

	return interface{}(dummyObj)
}

func printDetails(iObj interface{}) {
	switch iObj.(type) {
		case EndPoint:
			endPointObj := iObj.(EndPoint)
			fmt.Printf("type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
			fmt.Printf("type(endPointObj): %T,  value(iObj): %v\n", endPointObj, endPointObj)
			break
		case AccessPoint:
			accessPointObj := iObj.(AccessPoint)
			fmt.Printf("type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
			fmt.Printf("type(accessPointObj): %T,  value(iObj): %v\n", accessPointObj, accessPointObj)
			break
		default:
			fmt.Printf("type(iObj): %T,  value(iObj): %v\n", iObj, iObj)
			break
	}

	return
}


func main() {
	var i Creator

	/*
	objEndPoint := EndPoint{}
	i = objEndPoint
	dummyObjEndPoint := i.CreateObj()
	printDetails(dummyObjEndPoint)
	//fmt.Printf("dummyObjEndPoint: %#v\n", dummyObjEndPoint)
	*/

	objAccessPoint := AccessPoint{}
	i = objAccessPoint
	dummyObjAccessPoint := i.CreateObj()
	printDetails(dummyObjAccessPoint)
	//fmt.Printf("dummyObjAccessPoint: %#v\n", dummyObjAccessPoint)

	return
}
