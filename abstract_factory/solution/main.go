package solution

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

type VoucherAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}

type CoffeeMorningVoucherFactory struct{}

func (CoffeeMorningVoucherFactory) GetDrink() Drink {
	return Coffee{}
}

func (CoffeeMorningVoucherFactory) GetFood() Food {
	return Cake{}
}

type MidNightVoucherFactory struct{}

func (MidNightVoucherFactory) GetDrink() Drink {
	return Beer{}
}

func (MidNightVoucherFactory) GetFood() Food {
	return GrilledOctopus{}
}

func GetAbstractFactory(t string) VoucherAbstractFactory {
	switch t {
	case "mornign-coffee":
		return CoffeeMorningVoucherFactory{}
	case "chill-all-night":
		return MidNightVoucherFactory{}
	}
	return nil
}

func GetVoucher(t VoucherAbstractFactory) Voucher {
	return Voucher{
		Drink: t.GetDrink(),
		Food:  t.GetFood(),
	}
}
func main() {
	voucherAbs := GetAbstractFactory("mornign-coffee")
	myVoucher := GetVoucher(voucherAbs)
	fmt.Println(myVoucher)

}
