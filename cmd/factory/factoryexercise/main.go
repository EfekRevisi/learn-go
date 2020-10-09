package main

import (
	"fmt"
	factory "learning-go/pkg/factory/factoryexercise"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	maverick, _ := factory.GetGun("maverick")

	printDetails(ak47)
	printDetails(maverick)
}

func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
