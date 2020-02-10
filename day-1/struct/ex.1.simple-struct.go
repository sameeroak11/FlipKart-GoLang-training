package main

import (
	"fmt"
)


type device struct {
	ip string
	mac string
	mode int
}

type computer struct {
	bitStrength int
	core int
	os string
}

type dummy struct {
	id int
	tagname string
}

type infra struct {
	dev device
	devPC computer
}


func main() {
	dev := device {
		ip: "192.168.1.1",
		mac: "00:00:01",
		mode: 1,
	}
	fmt.Printf("dev: %#v\n", dev)

	pDev := &device {
		ip: "192.168.1.1",
		mac: "00:00:01",
		mode: 1,
	}
	fmt.Printf("pDev: %#v\n", *pDev)
	fmt.Printf("pDev.ip: %s\n", pDev.ip)
	fmt.Printf("pDev.ip: %s\n", pDev.mac)
	fmt.Printf("pDev.ip: %d\n\n", pDev.mode)

	instance1 := infra {
		dev: device {
			ip: "192.168.10.1",
			mac: "00:00:01",
			mode: 1,
		},
		devPC: computer {
			bitStrength: 64,
			core: 4,
			os: "ubuntu 18.04",
		},
	}

	fmt.Printf("infraStructure:\n")
	fmt.Printf("dev: %#v\n", instance1.dev)
	fmt.Printf("devPC: %#v\n", instance1.devPC)
}
