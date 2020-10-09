package interfacefactory

import "fmt"

type tiredperson struct {
	name string
	age  int
}

func (p *tiredperson) SayHello() {
	fmt.Printf("Sorry, I'm too tired")
}

func NewTiredPerson(name string, age int) Person {
	return &tiredperson{
		name,
		age,
	}
}
