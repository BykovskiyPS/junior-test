package FactoryMethod

import "testing"

func TestPizzaStore(t *testing.T) {
	assert := []string{"cheese", "salami"}

	pizzaStore := NewPizzaStore()
	pizzas := []Pizza{
		pizzaStore.CreatePizza("cheese"),
		pizzaStore.CreatePizza("salami"),
	}

	for i, pizza := range pizzas {
		if taste := pizza.Use(); taste != assert[i] {
			t.Errorf("Expect taste to %s, but %s.\n", assert[i], taste)
		}
	}
}
