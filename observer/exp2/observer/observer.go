package observer

import "fmt"

type Observer interface {
	updateItem(item string)
}

type Customer struct {
	Name string
}

func (c *Customer) updateItem(title string) {
	fmt.Printf("Woww!, %s will update and you can buy it now...Hurry up: %s\n", title, c.Name)
}
