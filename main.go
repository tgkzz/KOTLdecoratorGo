package main

import "fmt"

type Burger interface {
	getPrice() int
}

type SimpleBurger struct {
}

func (p *SimpleBurger) getPrice() int {
	return 25
}

type AddMeat struct {
	burger Burger
}

func (c *AddMeat) getPrice() int {
	AsdBurger := c.burger.getPrice()
	return AsdBurger + 35
}

type AddTomato struct {
	burger Burger
}

func (c *AddTomato) getPrice() int {
	AsdBurger := c.burger.getPrice()
	return AsdBurger + 25
}

type AddCheese struct {
	burger Burger
}

func (c *AddCheese) getPrice() int {
	AsdBurger := c.burger.getPrice()
	return AsdBurger + 15
}
func main() {
	burger := &SimpleBurger{}

	newburger := &AddTomato{
		burger: burger,
	}

	newnewburger := &AddMeat{
		burger: newburger,
	}

	newnewnewburger := &AddTomato{
		burger: newnewburger,
	}

	fmt.Println("the price of your burger is ", newnewnewburger.getPrice())
}
