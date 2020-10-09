package factoryexercise

import "fmt"

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}

	if gunType == "maverick" {
		return newMaverick(), nil
	}

	return nil, fmt.Errorf("Wrong gun type")
}
