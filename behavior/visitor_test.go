package behavior

import "testing"

func TestGetPriceCompute(t *testing.T) {
	products := make([]IVisitable, 2)
	products[0] = &OmegaSpeedMaster{
		Product: Product{
			Price: 7000.0,
			Name:  "Omega Speed Master",
		},
	}

	products[1] = &OmegaSeaMaster{
		Product: Product{
			Price: 5000.0,
			Name:  "Omega Sea Master",
		},
	}

	priceVisitor := &PriceVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
	}

	result := priceVisitor.Sum
	if result != 12000 {
		t.Errorf("must return  '%f'   not '%f' ", 12000.0, result)
	}

	t.Log(result)

}
