// slice of interfaces where each element of the slice has been assigned a
// reference of concrete-type variable where concrete type implements the interface.


package main

import (
	"fmt"
)

type DeviceModifier interface {
	CreateDevice()
	ModifyDevice()
	GetDeviceDetails()
	Dummy()
}

type Device struct {
	DevType string
	MAC string
	SrNo string
	DevMode string
}

type AccessPoint struct {
	DevType string
	WiFiProto string
	MAC string
	SrNo string
	Range uint8
	APMode string
}

func (pDev *Device) CreateDevice() {
	pDev.DevType = "0: EndPoint device"
	pDev.MAC = "0: abcd"
	pDev.SrNo = "0: s1"
	pDev.DevMode = "0: EndPoint"

	return
}
func (pDev *Device) ModifyDevice() {
	pDev.DevType = "1: Endpoint device"
	pDev.MAC = "1: abcd"
	pDev.SrNo = "1: s1"
	pDev.DevMode = "1: EndPoint mote"

	return
}
func (pDev *Device) GetDeviceDetails() {
	fmt.Printf("device type: %s,  device: %#v\n", pDev.DevType, pDev)
	return
}
func (dev Device) Dummy() {
	fmt.Printf("device calling Dummy()\n")
}


func (pAP *AccessPoint) CreateDevice() {
	pAP.DevType = "WiFi infrastructure device"
	pAP.WiFiProto = "0: 802.11 a/b/bg"
	pAP.MAC = "0: abcd"
	pAP.SrNo = "0: s1"
	pAP.Range = 50
	pAP.APMode = "0: AP"

	return
}
func (pAP *AccessPoint) ModifyDevice() {
	pAP.DevType = "WiFi infrastructure device"
	pAP.WiFiProto = "1: 802.11 a/b/bg/n"
	pAP.MAC = "1: abcd"
	pAP.SrNo = "1: s1"
	pAP.Range = 80
	pAP.APMode = "1: Repeater"

	return
}
func (pAP *AccessPoint) GetDeviceDetails() {
	fmt.Printf("device type: %s,  ap: %#v\n", pAP.DevType, pAP)
	return
}
func (ap AccessPoint) Dummy() {
	fmt.Printf("ap calling Dummy()\n")
}


func createDevices(ipDeviceList []DeviceModifier) {
	for _, pDev := range ipDeviceList {
		pDev.CreateDevice()
	}

	return
}

func updateDevices(ipDeviceList []DeviceModifier) {
	for _, pDev := range ipDeviceList {
		pDev.ModifyDevice()
	}

	return
}

func getDeviceDetails(msg string, ipDeviceList []DeviceModifier) {
	fmt.Printf("%s\n", msg)
	for _, piDev := range ipDeviceList {
		piDev.GetDeviceDetails()
	}

	//fmt.Printf("testing empty interface as receiver:\n")
	//&(AccessPoint{}).GetDeviceDetails()

	return
}


func main() {
	dev1 := Device{}
	ap1 := AccessPoint{}
	ipDeviceList := []DeviceModifier{&dev1, &ap1}

	/*
	the above assignment is as good as saying:
	var ip1 DeviceModifier
	var ip2 DeviceModifier

	ip1 = dev1
	ip2 = ap1

	*/

	createDevices(ipDeviceList)
	getDeviceDetails("creating devices:", ipDeviceList)

	fmt.Printf("\n")

	updateDevices(ipDeviceList)
	getDeviceDetails("updating devices:", ipDeviceList)

	dev1.Dummy()
	ap1.Dummy()

	return
}
