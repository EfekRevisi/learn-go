package main

import (
	factory "learning-go/pkg/factory/interfacefactory"
)

func main() {
	p := factory.NewNormalPerson("james", 34)
	p.SayHello()

	t := factory.NewTiredPerson("james", 34)
	t.SayHello()
}
