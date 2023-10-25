package structural

import (
	"strings"
	"testing"
)

func TestPizza_AddIngredient(t *testing.T) {
	// Onion decore notre Meat qui decore notre Pizza
	pizza := &Onion{&Meat{&Pizza{}}}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "ingredients: Meat, Onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("must return  '%s'   not '%s' ", expectedText, pizzaResult)
	}

	t.Log(pizzaResult)

}

func TestMeat_AddIngredient(t *testing.T) {
	// Meat decore notre Pizza
	pizza := &Meat{&Pizza{}}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "ingredients: Meat"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("must return  '%s'   not '%s' ", expectedText, pizzaResult)
	}

	t.Log(pizzaResult)

}

func TestOnion_AddIngredient(t *testing.T) {
	// Onion decore notre Pizza
	pizza := &Onion{&Pizza{}}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "ingredients: Onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("must return  '%s'   not '%s' ", expectedText, pizzaResult)
	}

	t.Log(pizzaResult)

}
