package internal

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (h *normalBuilder) setWindowType() {
	h.windowType = "normal window"
}

func (h *normalBuilder) setDoorType() {
	h.doorType = "normal door"
}

func (h *normalBuilder) setFloor() {
	h.floor = 6
}

func (h *normalBuilder) getHouse() House {
	return House{
		windowType: h.windowType,
		doorType:   h.doorType,
		floor:      h.floor,
	}
}
