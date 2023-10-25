package structural

import (
	"errors"
	"fmt"
)

// tous les decorators à venir implementent la même interface que le core
type IngredientAdd interface {
	AddIngredient() (string, error)
}

// core type Pizza
type Pizza struct {
	Ingredient IngredientAdd
}

func (p *Pizza) AddIngredient() (string, error) {
	return "ingredients:", nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("an ingredient is needed in Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "Meat"), nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("an ingredient is needed in Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "Onion"), nil
}
