package main

import (
	"fmt"
)


type device struct {
	ip string
	mac string  // this's key.
	mode int
}


func main() {
	var m1 = map[string]device {
		"00:00:00:00:00:01": device {
			ip: "192.168.1.1",
			mac: "00:00:00:00:00:01",
			mode: 1,
		},
		"00:00:00:00:00:02": device {
			ip: "192.168.1.2",
			mac: "00:00:00:00:00:02",
			mode: 2,
		},
	}

	fmt.Printf("m1[\"00:00:00:00:00:01\"]: %#v\n", m1["00:00:00:00:00:01"])
	fmt.Printf("m1[\"00:00:00:00:00:02\"]: %#v\n", m1["00:00:00:00:00:02"])

	fmt.Printf("\nm1:\n%#v\n", m1)
}
