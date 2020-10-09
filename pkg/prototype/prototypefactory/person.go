package prototypefactory

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Office Address
	Name   string
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{
	Name: "",
	Office: Address{
		Suite:         0,
		City:          "Bandung",
		StreetAddress: "Jalan Sukajadi no 14",
	},
}
var auxOffice = Employee{
	Name: "",
	Office: Address{
		Suite:         0,
		City:          "Padalarang",
		StreetAddress: "Jalan Cimareme no 1",
	},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite

	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}
