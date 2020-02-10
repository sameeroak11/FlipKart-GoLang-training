package main

import (
	"fmt"
	"reflect"
)


func f1(argList ...interface{}) {
	for _, rec := range argList {
		fmt.Println(rec, ":", reflect.ValueOf(rec).Kind())
	}
}


func main() {
	f1(1, "open bsd", true, 3.142857, []string{"ubuntu", "19.", "04"}, struct {id int
		mac string
		mode string
	}{id: 10, mac: "00:00:00:00:00:01", mode: "endpoint",},
		map[string]string{"os": "linux", "distro": "ubuntu", "version": "19.04",})
}
