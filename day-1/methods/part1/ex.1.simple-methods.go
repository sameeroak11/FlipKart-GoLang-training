package main

import (
	"fmt"
)


type employee struct {
	id int
	name string
}


type dummy struct {
	id int
	attr string
}

func (pDummy *dummy) Create(id int, attr string) {
	pDummy.id = 10
	pDummy.attr = attr
}

func (dummy dummy) Fetch() {
	fmt.Printf("dummy.id: %d\n", dummy.id)
	fmt.Printf("dummy.attr: %s\n", dummy.attr)
}


func (pEmp *employee) Create(id int, name string) {
	pEmp.id = id
	pEmp.name = name
}


func (emp employee) Fetch() {
	fmt.Printf("emp.id: %d\nemp.name: %s\n", emp.id, emp.name)
	fmt.Printf("emp: %#v\n", emp)
}


func main() {
	emp1 := employee{}
	emp1.Create(10, "abcd")
	emp1.Fetch()

	dummy := dummy{}
	dummy.Create(10, "attr1")
	dummy.Fetch()
}
