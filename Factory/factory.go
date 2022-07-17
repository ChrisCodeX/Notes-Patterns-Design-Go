package main

type iProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type computer struct {
	name  string
	stock int
}

func (c computer) setStock(stock int) {
	c.stock = stock
}

func (c computer) setName(name string) {
	c.name = name
}

func (c computer) getStock() int {
	return c.stock
}

func (c computer) getName() string {
	return c.name
}

// Laptop
type laptop struct {
	computer
}

func newLaptop() iProduct {
	return &laptop{
		computer: computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

// Desktop
type desktop struct {
	computer
}
