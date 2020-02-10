package main

import (
	"fmt"
)


func main() {
	m1 := make(map[string]int)

	m1["one"] = 1
	m1["two"] = 2
	m1["ten"] = 10
	m1["twelve"] = 12

	fmt.Printf("m1: %#v\n", m1)

	fmt.Printf("m1[\"one\"]: %d\n", m1["one"])
	fmt.Printf("m1[\"two\"]: %d\n", m1["two"])

	fmt.Printf(`m1["one"]: %d%s`, m1["one"], "\n")
	fmt.Printf(`m1["two"]: %d%s`, m1["two"], "\n")

	if val, isOK := m1["one"]; isOK {
		fmt.Printf("key exists. val: %d\n", val)
	}

	if val, isOK := m1["three"]; !isOK {
		fmt.Printf("key doesn't exist. val: %d\n", val)
	}

	for key, val := range m1 {
		fmt.Printf("key:val in range: %s.%d\n", key, val)
	}

	for _, val := range m1 {
		fmt.Printf("key ingnored. val in range: %d\n", val)
	}

	for key, _ := range m1 {
		fmt.Printf("val ingnored. key in range: %s\n", key)
	}

	for key := range m1 {
		fmt.Printf("val ingnored again. key in range: %s\n", key)
	}

	delete(m1, "one")
	for key, val := range m1 {
		fmt.Printf("after delete one: key:val in range: %s.%d\n", key, val)
	}

	for key, val := range m1 {
		val = 10
		m1[key] = val
		fmt.Printf("mutation: delete one: key:val in range: %s.%d\n", key, val)
	}
}
