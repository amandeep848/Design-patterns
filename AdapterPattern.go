package main

import "fmt"

// Target
type mobile interface {
	chargeAppleMobile()
}

// Concrete proto type implementation
type apple struct{}

func (a *apple) chargeAppleMobile() {
	fmt.Print("Charging Apple mobile")
}

// Client
type client struct{}

func (c *client) chargeMobile(m mobile) {
	c.chargeMobile(m)
}
func main() {

	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)

}

// When you don't want to change existing object or interface rather want to add new functionality on top of existing one
