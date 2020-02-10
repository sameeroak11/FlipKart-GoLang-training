package main


import (
	"fmt"
)


type DeviceModifier interface {
	CreateDevice()
	ModifyDevice()
	GetDevice()
	Foobar()
}


type Device struct {
	DevType string
	MAC string
	SrNo string
}

type Computer struct {
	DevType string
	Proc string
	MAC string
	SrNo string
}

func (dev Device) Foobar() {
	fmt.Printf("dev Foobar()\n")
}
func (dev Device) CreateDevice() {
	dev.DevType = "0: Endpoint"
	dev.MAC = "0: abcd"
	dev.SrNo = "0: s1"
	dev.GetDevice()
}
func (dev Device) ModifyDevice() {
	dev.DevType = "1: Endpoint"
	dev.MAC = "1: abcd"
	dev.SrNo = "1: s1"

	dev.GetDevice()
}
func (dev Device) GetDevice() {
	fmt.Printf("dev-type: %s,  device: %#v\n", dev.DevType, dev)
}

// computer type:
func (comp Computer) Foobar() {
	fmt.Printf("comp Foobar()\n")
}
func (dev Computer) CreateDevice() {
	dev.DevType = "0: Computer"
	dev.Proc = "0: x86_64"
	dev.MAC = "0: abcd"
	dev.SrNo = "0: s1"

	dev.GetDevice()
}
func (dev Computer) ModifyDevice() {
	dev.DevType = "1: Computer"
	dev.Proc = "1: amd 86_64"
	dev.MAC = "1: abcd"
	dev.SrNo = "1: s1"

	dev.GetDevice()
}
func (dev Computer) GetDevice() {
	fmt.Printf("dev-type: %s,  device: %#v\n", dev.DevType, dev)
}

func main() {
	dev := Device{}
	comp := Computer{}

	var dm DeviceModifier

	dm = dev
	dm.CreateDevice()
	dm.ModifyDevice()

	dm = comp
	dm.CreateDevice()
	dm.ModifyDevice()

	return
}
