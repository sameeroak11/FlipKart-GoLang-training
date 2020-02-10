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
	CreateObj()
}

type Mutator interface {
	MutateObj()
}

type ObjOperator interface {
	Creator
	Mutator
	Fetchdetails(string)
}

func (pEndPoint *EndPoint) CreateObj() {
	*pEndPoint = EndPoint {
		MAC: "[0: EndPoint]: abcd",
		SrNo: "[0: EndPoint]: 012",
	}

	return
}

func (pEndPoint *EndPoint) MutateObj() {
	*pEndPoint = EndPoint {
		MAC: "[1: EndPoint]: xyz",
		SrNo: "[1: EndPoint]: 987",
	}

	return
}

func (pEndPoint *EndPoint) Fetchdetails(msg string) {
	fmt.Printf("[%s]: type(iObj): %T,  value(iObj): %v\n", msg, pEndPoint, pEndPoint)
}


func main() {
	var i ObjOperator

	objEndPoint := EndPoint{}
	i = &objEndPoint
	i.CreateObj()
	i.Fetchdetails("creator")
	i.MutateObj()
	i.Fetchdetails("mutator")

	return
}
