// methods:
// methods with pointer receiver and on userdefined type on built-in data-type.


package main

import (
	"fmt"
)

type DevType string

func (devType *DevType) dummy() {
	fmt.Printf("in dummy: dev-type: %s\n", *devType)
	*devType = "abcd"
}

func (devType *DevType) GetDevType() {
	fmt.Printf("dev-type: %s\n", *devType)
}

func (devType *DevType) SetDevType(newType DevType) {
	//devType = "WiFi Endpoint Mote"
	*devType = newType
}

func main() {
	devType := DevType("WiFi Endpoint")
	(&devType).GetDevType()		// WiFi Endpoint
	//(&devType).SetDevType(DevType("WiFi Endpoint Mote"))
	(&devType).SetDevType("WiFi Endpoint Mote")
	(&devType).GetDevType()		// WiFi Endpoint Mote

	devType.dummy()
	fmt.Printf("in main: dev-type: %s\n", devType)

	return
}
