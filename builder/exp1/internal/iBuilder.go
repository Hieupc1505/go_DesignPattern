package internal

type IBuilder interface {
	setWindowType()
	setDoorType()
	setFloor()
	getHouse() House
}

func NewIBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}

	}
	return nil
}
