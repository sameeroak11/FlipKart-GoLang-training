// methods:
// pointer receiver


package main

import (
	"fmt"
)

type Device struct {
	IP string
	MAC string
}

type Computer struct {
	Processor string
	RAM int8
	DiskSize int16
}

func Create(partID int) {
	fmt.Printf("partID: %d\n", partID)
}


func (pDev **Device) Create(ip string, mac string) {
	fmt.Printf("type of pDev: %T\n", pDev)
	pDev.IP = ip
	pDev.MAC = mac

	return
}

/* func (dev Device) Create(ip string, mac string) Device {
	dev.IP = ip
	dev.MAC = mac

	fmt.Printf("ip: %s, mac: %s\n", ip, mac)

	return dev
} */

func (pDev **Computer) Create(proc string, ram int8, disksize int16) {
	fmt.Printf("type of pDev: %T\n", pDev)
	pDev.Processor = "abcd"
	pDev.RAM = 8
	pDev.DiskSize = 500

	return
}

func main() {
	dev := Device{}
	comp := Computer{}

	//(&dev).Create("192.168.1.10", "abcd")
	//(&comp).Create(true, &comp, "quad core", 16, 1000)
	pDev := &dev
	pComp := &comp

	//dev.Create("192.168.1.10", "abcd")
	//comp.Create(true, "quad core", 16, 1000)
	&pDev.Create("192.168.1.10", "abcd")
	&pComp.Create("quad core", 16, 1000)

	fmt.Printf("dev: %#v\n", dev)
	fmt.Printf("comp: %#v\n", comp)

	Create(10)

	return
}
