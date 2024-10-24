package extend

import "fmt"

type IChair interface {
	setName(name string)
	getName() string
}

type Chair struct {
	name string
}

func (c *Chair) setName(name string) {
	c.name = name
}

func (c *Chair) getName() string {
	return c.name
}

type ITable interface {
	setName(name string)
	getName() string
}

type Table struct {
	name string
}

func (t *Table) setName(name string) {
	t.name = name
}

func (t *Table) getName() string {
	return t.name
}

type ISofa interface {
	setName(name string)
	getName() string
}

type Sofa struct {
	name string
}

func (s *Sofa) setName(name string) {
	s.name = name
}

func (s *Sofa) getName() string {
	return s.name
}

type IProductFactory interface {
	GetInfo()
	SetInfo(chair IChair, table ITable, sofa ISofa)
}

type ProductFactory struct {
	chair IChair
	table ITable
	sofa  ISofa
}

func (a *ProductFactory) GetInfo() {
	fmt.Println("My products are: %s - %s - %s", a.chair.getName(), a.table.getName(), a.sofa.getName())
}

func (a *ProductFactory) SetInfo(chair IChair, table ITable, sofa ISofa) {
	a.chair = chair
	a.table = table
	a.sofa = sofa
}

func main() {
	artDecoChair := &Chair{
		name: "Art Deco Chair",
	}

	artDecoSofa := &Sofa{
		name: "Art Deco Sofa",
	}

	artDecoTable := &Table{
		name: "Art Deco Table",
	}

	productFactory := &ProductFactory{
		chair: artDecoChair,
		table: artDecoTable,
		sofa:  artDecoSofa,
	}
	productFactory.GetInfo()
}
