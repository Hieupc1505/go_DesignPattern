package yuh

type IShoe interface {
	setLogo(logo string)
	GetLogo() string
	setSize(size int)
	GetSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) GetLogo() string {
	return s.logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) GetSize() int {
	return s.size
}

type ISkirt interface {
	setLogo(logo string)
	GetLogo() string
	setSize(size int)
	GetSize() int
}

type skirt struct {
	logo string
	size int
}

func (s *skirt) setLogo(logo string) {
	s.logo = logo
}

func (s *skirt) GetLogo() string {
	return s.logo
}

func (s *skirt) setSize(size int) {
	s.size = size
}

func (s *skirt) GetSize() int {
	return s.size
}

/*
	Nếu như shirt model trên được tách vào một packeage thì sẽ không có cách nào để
	khởi tạo một đối tượng skirt vì shirt là private ngay cả khi sử dụng
	model.shirt{logo: 'adidas', size: 20} thì cũng không được, do đó nếu
	muốn giữ cách trên thì phải export một contructor khởi tạo đối tượng
	đối tượng shirt ra bên ngoài ví dụ:
	func CreateShirt(logo string, size number) IShirt {
		return &shirt{
			logo: logo,
			size: size,
		}
	}
*/

type AbstractFactory interface {
	CreateShoe() IShoe
	CreateSkirt() ISkirt
}

type NikeFactory struct{}

func (NikeFactory) CreateShoe() IShoe {
	return &shoe{
		logo: "shoe nike logo",
		size: 10,
	}
}

func (NikeFactory) CreateSkirt() ISkirt {
	return &shoe{
		logo: "skirt nike logo",
		size: 10,
	}
}

type adidasFactory struct{}

func (adidasFactory) CreateShoe() IShoe {
	return &shoe{
		logo: "shoe adidas logo",
		size: 10,
	}
}

func (adidasFactory) CreateSkirt() ISkirt {
	return &shoe{
		logo: "skirt adidas logo",
		size: 10,
	}
}

func GetAbstractFactory(brand string) AbstractFactory {
	switch brand {
	case "adidas":
		return adidasFactory{}
	case "nike":
		return NikeFactory{}
	default:
		return adidasFactory{}
	}
}

func main() {
	factory := GetAbstractFactory("adidas")
	shoe := factory.CreateShoe()
	skirt := factory.CreateSkirt()

	println(shoe.GetLogo())
	println(shoe.GetSize())
	println(skirt.GetLogo())
	println(skirt.GetSize())
}
