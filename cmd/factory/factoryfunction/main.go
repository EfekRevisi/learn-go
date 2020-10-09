package main

import (
	"fmt"
	factory "learning-go/pkg/factory/factoryfunction"
)

func main() {
	p := factory.Person{"John", 22}
	fmt.Println(p)

	p2 := factory.NewPerson("John 2", 23)
	p2.Age = 20
	fmt.Println(p2)
}
