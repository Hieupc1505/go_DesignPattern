package main

import "fmt"

/*
Issue: Đúng vậy, khi triển khai factory cho từng sản phẩm cụ thể như trong đoạn mã của bạn, bạn sẽ gặp vấn đề lặp lại mã.
Để giảm thiểu việc lặp lại này, bạn có thể tạo một factory chung cho tất cả các sản phẩm mà vẫn giữ nguyên tính linh hoạt và khả năng mở rộng.

Giải pháp: Sử dụng một Factory Chung
Bạn có thể định nghĩa một factory chung cho các sản phẩm và tạo ra các sản phẩm dựa trên tham số đầu vào. Dưới đây là cách bạn có thể làm điều này:
*/

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

// Interface cho Factory
type IProductFactory interface {
	CreateChair() IChair
	CreateTable() ITable
	CreateSofa() ISofa
}

// Factory cụ thể cho sản phẩm Art Deco
type ArtDecoFactory struct{}

func (f *ArtDecoFactory) CreateChair() IChair {
	return &Chair{name: "Art Deco Chair"}
}

func (f *ArtDecoFactory) CreateTable() ITable {
	return &Table{name: "Art Deco Table"}
}

func (f *ArtDecoFactory) CreateSofa() ISofa {
	return &Sofa{name: "Art Deco Sofa"}
}

// Factory cho sản phẩm
type ProductFactory struct {
	chair IChair
	table ITable
	sofa  ISofa
}

func (a *ProductFactory) GetInfo() {
	fmt.Printf("My products are: %s - %s - %s\n", a.chair.getName(), a.table.getName(), a.sofa.getName())
}

func main() {
	// Tạo factory cụ thể
	factory := &ArtDecoFactory{}

	// Tạo sản phẩm thông qua factory
	productFactory := &ProductFactory{
		chair: factory.CreateChair(),
		table: factory.CreateTable(),
		sofa:  factory.CreateSofa(),
	}
	productFactory.GetInfo()
}
