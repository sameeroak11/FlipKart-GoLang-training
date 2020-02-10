// interfaces:
// methods with pointer receiver


package main

import (
	"fmt"
)

type DeviceModifier interface {
	CreateDevice()
	ModifyDevice()
	DeviceDetails(string)
}

type Device struct {
	DevType string
	MAC string
	SrNo string
	DevMode string
}

func (pDev *Device) CreateDevice() {
	pDev.DevType = "EndPoint device"
	pDev.MAC = "0: abcd"
	pDev.SrNo = "0: s1"
	pDev.DevMode = "0: EndPoint"
	pDev.DeviceDetails("create")
}

func (pDev *Device) ModifyDevice() {
	pDev.DevType = "Endpoint device"
	pDev.MAC = "1: abcd"
	pDev.SrNo = "1: s1"
	pDev.DevMode = "1: EndPoint mote"
	pDev.DeviceDetails("modify")
}

func (pDev *Device) DeviceDetails(msg string) {
	fmt.Printf("[%s]: device-type: %s, device-details: %#v\n", msg, pDev.DevType, pDev)
}

func main() {
	var i DeviceModifier
	dev := Device{}
	//(&dev).CreateDevice()
	i = &dev
	i.CreateDevice()
	i.ModifyDevice()

	return
}
