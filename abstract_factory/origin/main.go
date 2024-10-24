package main

import (
	"fmt"
)

// iSportsFactory.go: Abstract factory interface
type iSportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (iSportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("Wrong brand type passed")
	}
}

// adidas.go: Concrete factory
type Adidas struct{}

func (s *Adidas) makeShoe() IShoe {
	return &Shoe{
		name: "adidas shoe",
		size: 14,
	}
}

func (s *Adidas) makeShirt() IShirt {
	return &Shirt{
		name:  "adidas shirt",
		size:  20,
		color: "blue",
	}
}

// nike.go: Concrete factory
type Nike struct{}

func (s *Nike) makeShoe() IShoe {
	return &Shoe{
		name: "nike shoe",
		size: 12,
	}
}

func (s *Nike) makeShirt() IShirt {
	return &Shirt{
		name:  "nike shirt",
		size:  18,
		color: "white",
	}
}

// iShoe.go: Abstract product
type IShoe interface {
	setName(string)
	getName() string
	setSize(int)
	getSize() int
}

type Shoe struct {
	name string
	size int
}

func (s *Shoe) setName(name string) {
	s.name = name
}

func (s *Shoe) getName() string {
	return s.name
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}

// adidasShoe.go: Concrete product
type AdidasShoe struct {
	Shoe
}

// nikeShoe.go: Concrete product
type NikeShoe struct {
	Shoe
}

// iShirt.go: Abstract product
type IShirt interface {
	setName(string)
	getName() string
	setSize(int)
	getSize() int
	setColor(string)
	getColor() string
}

type Shirt struct {
	name  string
	size  int
	color string
}

func (s *Shirt) setName(name string) {
	s.name = name
}

func (s *Shirt) getName() string {
	return s.name
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}

func (s *Shirt) setColor(color string) {
	s.color = color
}

func (s *Shirt) getColor() string {
	return s.color
}

// adidasShirt.go: Concrete product
type AdidasShirt struct {
	Shirt
}

// nikeShirt.go: Concrete product
type NikeShirt struct {
	Shirt
}

// make product
type SportProduct struct {
	Shoe  IShoe
	Shirt IShirt
}

func getSportProduct(factory iSportsFactory) SportProduct {
	return SportProduct{
		Shoe:  factory.makeShoe(),
		Shirt: factory.makeShirt(),
	}
}

// main.go: Client code
func main() {
	Adidasfactory, _ := GetSportsFactory("adidas")
	product := getSportProduct(Adidasfactory)

	fmt.Printf("My adidas's product is: %v\n%v", product.Shoe, product.Shirt)

	Nikefactory, _ := GetSportsFactory("nike")
	product = getSportProduct(Nikefactory)

	fmt.Println("My nike's product is: ", product)
}
