package builderexercise

type normalHouse struct {
	windowType, doorType string
	floor                int
}

func NewNormalHouseBuilder() *normalHouse {
	return &normalHouse{}
}

func (b *normalHouse) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalHouse) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalHouse) setNumFloor() {
	b.floor = 2
}

func (b *normalHouse) getHouse() house {
	return house{
		DoorType:   b.doorType,
		Floor:      b.floor,
		WindowType: b.windowType,
	}
}
