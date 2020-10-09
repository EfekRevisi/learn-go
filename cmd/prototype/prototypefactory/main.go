package main

import (
	"fmt"
	"learning-go/pkg/prototype/prototypefactory"
)

func main() {
	john := prototypefactory.NewMainOfficeEmployee("John", 100)
	jane := prototypefactory.NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
