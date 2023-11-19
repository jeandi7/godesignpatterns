package main

import "fmt"

// MonoidElement represents the interface for elements in a monoid.
type MonoidElement interface {
	Compose(element MonoidElement) MonoidElement
	NeutralElement() MonoidElement
}

// IntegerMonoidElement is an implementation of MonoidElement for integers with addition as the composition operation.
type IntegerMonoidElement int

func (i IntegerMonoidElement) Compose(element MonoidElement) MonoidElement {
	// The composition operation for integers is addition.
	return IntegerMonoidElement(int(i) + int(element.(IntegerMonoidElement)))
}

func (i IntegerMonoidElement) NeutralElement() MonoidElement {
	// The neutral element for addition is 0.
	return IntegerMonoidElement(0)
}

func main() {
	// Example usage of the generic MonoidElement interface.

	// Creating instances of MonoidElement for integers.
	monoid1 := IntegerMonoidElement(5)
	monoid2 := IntegerMonoidElement(10)

	// Composing the monoid elements.
	result := monoid1.Compose(monoid2)

	// Displaying the result.
	fmt.Println(result) // Output: 15

	// Checking associativity.
	monoid3 := IntegerMonoidElement(8)
	associativityCheck := monoid1.Compose(monoid2).Compose(monoid3)
	associativityCheck2 := monoid1.Compose(monoid2.Compose(monoid3))

	// The two results should be equal if associativity is respected.
	fmt.Println(associativityCheck == associativityCheck2) // Output: true
}
