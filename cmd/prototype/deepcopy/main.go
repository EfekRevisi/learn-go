package main

import (
	"fmt"
	"learning-go/pkg/prototype/deepcopy"
)

func main() {
	john := deepcopy.Person{
		Name: "John",
		Address: &deepcopy.Address{
			City:          "Padalarang",
			Country:       "ID",
			StreetAddress: "Jln Cimareme",
		},
	}

	// shallow copy
	// jane := john

	// deep copy
	jane := john
	jane.Address = &deepcopy.Address{
		City:          john.Address.City,
		Country:       john.Address.Country,
		StreetAddress: john.Address.StreetAddress,
	}

	jane.Name = "Jane"
	jane.Address.StreetAddress = "Jln Ciburuy"

	fmt.Println(john.Name, john.Address)
	fmt.Println(jane.Name, jane.Address)
}
