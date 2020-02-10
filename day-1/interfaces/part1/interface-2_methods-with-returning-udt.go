// interfaces:
// methods returning userdefined type


package main


import (
	"fmt"
)


type DeviceModifier interface {
	CreateDevice() Device
	ModifyDevice() Device
}


type Device struct {
	MAC string
	SrNo string
}

type Computer struct {
	Proc string
	MAC string
	SrNo string
}

func (dev Device) CreateDevice() Device {
	dev.MAC = "0: abcd"
	dev.SrNo = "0: s1"

	return dev
}

func (dev Device) ModifyDevice() Device {
	dev.MAC = "1: abcd"
	dev.SrNo = "1: s1"

	return dev
}

func main() {
	dev := Device{}
	var dm DeviceModifier

	dm = dev
	fmt.Printf("creating device-1: %#v\n", dm.CreateDevice())
	fmt.Printf("updating device-1: %#v\n", dm.ModifyDevice())

	return
}
