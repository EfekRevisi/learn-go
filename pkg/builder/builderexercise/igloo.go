package builderexercise

type igloo struct {
	windowType, doorType string
	floor                int
}

func NewIglooBuilder() *igloo {
	return &igloo{}
}

func (b *igloo) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *igloo) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *igloo) setNumFloor() {
	b.floor = 1
}

func (b *igloo) getHouse() house {
	return house{
		DoorType:   b.doorType,
		Floor:      b.floor,
		WindowType: b.windowType,
	}
}
