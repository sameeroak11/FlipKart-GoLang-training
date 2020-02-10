package main


import (
	"fmt"
)


func main() {
	intArray1 := [4]int {10, 20, 30, 40,}
	intArray2 := [...]int {10, 20, 30, 40}

	intArray3 := [2][4]int {
		{10, 20, 30, 40},
		{100, 200, 300, 400},
	}

	intArray4 := [...][4]int {{1, 2, 3, 4,},
		{5, 6, 7, 8}}

	fmt.Println("intArray1:", intArray1)
	fmt.Println("intArray2:", intArray2)
	fmt.Println("intArray3:", intArray3)
	fmt.Println("intArray4:", intArray4)

	var intArray5 [4]int
	intArray5[0] = 1000
	intArray5[1] = 2000
	intArray5[2] = 3000
	intArray5[3] = 4000
	fmt.Println("intArray5:", intArray5)


	fmt.Printf("intArray5: %#v\n", intArray5)

	type dev struct {
		ip string
		mac string
		ssid string
	}

	deviceType := [2]dev {
		{
			mac: "aldaldf",
			ssid: "dlahdlsha",
		},
		{
			ip: "ldhldsha",
			mac: "aldaldf",
			ssid: "dlahdlsha",
		},
	}

	//emptyInterface := [4]interface{}{10, "abcd", 1.20, false}
}
