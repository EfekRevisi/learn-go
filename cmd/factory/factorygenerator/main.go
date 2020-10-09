package main

import (
	"fmt"
	factory "learning-go/pkg/factory/factorygenerator"
)

func main() {
	developerFactory := factory.NewEmployeeFactory("Engineer", 1000000)
	managerFactory := factory.NewEmployeeFactory("Engineer", 1800000)

	developer := developerFactory("Irfan")
	fmt.Println(developer)

	manager := managerFactory("John")
	fmt.Println(manager)

	bossFactory := factory.NewEmployeeFactory2("Boss", 3000000)
	bossFactory.AnnualIncome = 3500000
	boss := bossFactory.SetName("James")
	fmt.Println(boss)
}
