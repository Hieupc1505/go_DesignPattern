package internal

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (h *iglooBuilder) setWindowType() {
	h.windowType = "Igloo Window"
}

func (h *iglooBuilder) setDoorType() {
	h.doorType = "Igloo Door"
}

func (h *iglooBuilder) setFloor() {
	h.floor = 4
}

func (h *iglooBuilder) getHouse() House {
	return House{
		windowType: h.windowType,
		doorType:   h.doorType,
		floor:      h.floor,
	}
}
