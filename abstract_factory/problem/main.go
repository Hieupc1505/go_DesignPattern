package problem

import "fmt"

type Drink interface {
	Drink()
}

type Food interface {
	Eat()
}

type Voucher struct {
	Drink
	Food
}

type Coffee struct{}

func (Coffee) Drink() {
	fmt.Println("It's coffee, drinkable")
}

type Beer struct{}

func (Beer) Drink() {
	fmt.Println("Its beer, drinkable")
}

type Cake struct{}

func (Cake) Eat() {
	fmt.Println("Its cake, eatable")
}

type GrilledOctopus struct{}

func (GrilledOctopus) Eat() {
	fmt.Println("Its octopus, eatable")
}

func main() {
	fmt.Println(
		[]Voucher{
			//Voucher with Coffee & Cake, its good
			Voucher{
				Drink: Coffee{},
				Food:  Cake{},
			},
			//This voucher is great!!
			Voucher{
				Drink: Beer{},
				Food:  GrilledOctopus{},
			},
			//Oh, this voucher quite wired, I'm not sure what to do
			Voucher{
				Drink: Coffee{},
				Food:  GrilledOctopus{},
			},
		})

}
