package builderexercise

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return &normalHouse{}
	}

	if builderType == "igloo" {
		return &igloo{}
	}

	return nil
}
