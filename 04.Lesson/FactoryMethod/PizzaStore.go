package FactoryMethod

import (
	"fmt"
)

type PizzaStore interface {
	CreatePizza(taste string) Pizza
}

type Pizza interface {
	Use() string
}

type MoscowPizzaStore struct {
}

func NewPizzaStore() PizzaStore {
	return &MoscowPizzaStore{}
}

func (f *MoscowPizzaStore) CreatePizza(taste string) Pizza {
	var pizza Pizza

	switch taste {
	case "cheese":
		pizza = &CheesePizza{taste}
	case "salami":
		pizza = &PepperoniPizza{taste}
	default:
		fmt.Println("Unknown taste")
	}

	return pizza
}

type CheesePizza struct {
	taste string
}

func (c *CheesePizza) Use() string {
	return c.taste
}

type PepperoniPizza struct {
	taste string
}

func (p *PepperoniPizza) Use() string {
	return p.taste
}
