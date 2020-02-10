// methods:
// Same function name with varied receiver type and varied argument list.

/* func (receiver_name receiver_type) function_name([arguments...]) [(return_parameter)...] {
} */


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


/*func (pDev *Device) Create(ip string, mac string) {
	fmt.Printf("type of pDev: %T\n", pDev)
	pDev.IP = ip
	pDev.MAC = mac

	return
}*/

func (dev Device) Create(ip string, mac string) Device {
	dev.IP = ip
	dev.MAC = mac

	fmt.Printf("ip: %s, mac: %s\n", ip, mac)

	return dev
}

func (pDev *Computer) Create(dirFalg bool, proc string, ram int8, disksize int16) {
	fmt.Printf("type of pDev: %T\n", pDev)
	pDev.Processor = "abcd"
	pDev.RAM = 8
	pDev.DiskSize = 500

	/*
	if !dirFalg {
		return pDev
	}

	pDev.Processor = proc
	pDev.RAM = ram
	pDev.DiskSize = disksize
	*/

	return //pDev1
}

func (pDev *Computer) GetComputer() {
	fmt.Printf("before:  computer.processor: %s,  computer.RAM: %d, computer.disksize: %d\n", pDev.Processor, pDev.RAM, pDev.DiskSize)
	pDev.RAM = pDev.RAM + 1
	fmt.Printf("after:  computer.processor: %s,  computer.RAM: %d, computer.disksize: %d\n", pDev.Processor, pDev.RAM, pDev.DiskSize)
}

func main() {
	//dev := Device{}
	comp := Computer{}

	//(&dev).Create("192.168.1.10", "abcd")
	//(&comp).Create(true, &comp, "quad core", 16, 1000)
	//dev.Create("192.168.1.10", "abcd")
	comp.Create(true, "quad core", 16, 1000)
	comp.GetComputer()
	fmt.Printf("in main, dev: %#v\n", Device{}.Create("192.168.1.20", "abcd"))

	//fmt.Printf("dev: %#v\n", dev)
	fmt.Printf("comp: %#v\n", comp)
	//fmt.Printf("dev1: %#v\n", 
	//fmt.Printf("dev1: %#v\n", dev.Create("192.168.1.20", "abcd"))
	//fmt.Printf("comp-1: %#v\n", comp)

	Create(10)

	return
}
