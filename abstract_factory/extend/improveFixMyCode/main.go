package main

import "fmt"

// Interface cho các sản phẩm
type IChair interface {
	setName(name string)
	getName() string
}

type ITable interface {
	setName(name string)
	getName() string
}

type ISofa interface {
	setName(name string)
	getName() string
}

// Các sản phẩm cụ thể
type Chair struct {
	name string
}

func (c *Chair) setName(name string) {
	c.name = name
}

func (c *Chair) getName() string {
	return c.name
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

type Sofa struct {
	name string
}

func (s *Sofa) setName(name string) {
	s.name = name
}

func (s *Sofa) getName() string {
	return s.name
}

// Factory chung cho các sản phẩm
type ProductFactory struct{}

func (f *ProductFactory) CreateChair(style string) IChair {
	return &Chair{name: style + " Chair"}
}

func (f *ProductFactory) CreateTable(style string) ITable {
	return &Table{name: style + " Table"}
}

func (f *ProductFactory) CreateSofa(style string) ISofa {
	return &Sofa{name: style + " Sofa"}
}

// Factory cho sản phẩm
type Factory struct {
	chair IChair
	table ITable
	sofa  ISofa
}

func (a *Factory) GetInfo() {
	fmt.Printf("My products are: %s - %s - %s\n", a.chair.getName(), a.table.getName(), a.sofa.getName())
}

func main() {
	factory := &ProductFactory{}

	// Tạo các sản phẩm thông qua factory
	productFactory := &Factory{
		chair: factory.CreateChair("Art Deco"),
		table: factory.CreateTable("Art Deco"),
		sofa:  factory.CreateSofa("Art Deco"),
	}
	productFactory.GetInfo()
}

/*
Giải thích:
ProductFactory: Factory này có ba phương thức (CreateChair, CreateTable, CreateSofa)
nhưng giờ đây bạn có thể truyền tham số style để tạo ra các sản phẩm với các phong cách khác nhau.

Tránh lặp mã: Với cách triển khai này, bạn không cần tạo các factory khác nhau cho từng phong cách.
Bạn chỉ cần gọi CreateChair, CreateTable, hoặc CreateSofa với phong cách mong muốn.

Dễ dàng mở rộng: Bạn có thể dễ dàng thêm phong cách mới mà không cần thay đổi nhiều trong mã nguồn.
Chỉ cần gọi factory với phong cách khác nhau.

Kết luận:
Cách tiếp cận này giúp bạn tổ chức mã một cách gọn gàng hơn, giảm thiểu sự lặp lại, và vẫn duy trì tính linh hoạt trong việc tạo ra các sản phẩm khác nhau.
*/
