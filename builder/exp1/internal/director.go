package internal

type DirectorService interface {
	NewDirector(b IBuilder) *Director
}

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func SetBuilder(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (h *Director) BuildHouse() House {
	h.builder.setWindowType()
	h.builder.setDoorType()
	h.builder.setFloor()

	return h.builder.getHouse()
}
