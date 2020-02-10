package main

import (
	"fmt"
)


type device struct {
	ip string
	mac string  // this's key.
	mode int
}

var m1 map[string]device

func main() {
	m1 = make(map[string]device)
	dev := device{}
	/* m1["00:00:00:00:00:01"] = device {
		mac: "00:00:00:00:00:01",
		mode: 1,
		ip: "192.168.1.1",
	} */
	m1["00:00:00:00:00:01"] = dev

	fmt.Printf("m1[\"00:00:00:00:00:01\"]: %#v\n", m1["00:00:00:00:00:01"])

	fmt.Println()
	m2 := make(map[string]*device)
	pDev := &device{}
	m2["abcd"] = pDev
	pTmp := m2["abcd"]
	pTmp.mode = 1000
	fmt.Printf("m2[\"abcd\"]: %#v\n", m2["abcd"])
}
