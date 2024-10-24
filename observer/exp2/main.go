package main

import ob "github.com/hieupc05/goDesignPattern/observer/exp2/observer"

func main() {

	ip15 := ob.NewItem("IP 15 promax")
	ip16 := ob.NewItem("IP 16")
	cust1 := &ob.Customer{
		Name: "Hieupc",
	}

	cust2 := &ob.Customer{
		Name: "hieucc2",
	}
	ip15.Register(cust1)
	ip15.Register(cust2)
	ip16.Register(cust1)

	ip15.UpdateStockItem()
	ip16.UpdateStockItem()

}
