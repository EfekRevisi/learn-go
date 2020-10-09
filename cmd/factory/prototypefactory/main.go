package main

import (
	"fmt"
	prototype "learning-go/pkg/factory/prototypefactory"
)

func main() {
	m := prototype.NewEmployee(prototype.Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
