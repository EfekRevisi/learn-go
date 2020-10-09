package main

import (
	"fmt"
	"learning-go/pkg/prototype/serialization"
)

func main() {
	john := serialization.Person{
		Name: "John",
		Address: &serialization.Address{
			City:          "Padalarang",
			Country:       "ID",
			StreetAddress: "Jln Cimareme",
		},
		Friends: []string{
			"Chriss", "Patt",
		},
	}

	// shallow copy
	// jane := john

	// deep copy
	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "Jln Ciburuy"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
