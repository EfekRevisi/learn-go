package main

import (
	"fmt"
	"learning-go/pkg/prototype/copymethod"
)

func main() {
	john := copymethod.Person{
		Name: "John",
		Address: &copymethod.Address{
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
