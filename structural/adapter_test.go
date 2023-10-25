package structural

import (
	"testing"
)

func TestAdapter_Paiement(t *testing.T) {

	paymentGateway := &PayPalAdapter{
		PayPal: &PayPal{},
	}

	paymentGateway2 := &AmazonAdapter{
		Amazon: &Amazon{},
	}

	// triggers PayPal
	payR := paymentGateway.ProcessPayment(100)

	// triggers Amazon
	payR2 := paymentGateway2.ProcessPayment(100)

	if payR != true {
		t.Error("the payment Paypal must be true but it is \n", payR)
	}

	if payR2 != true {
		t.Error("the payment Amazon  must be true but it is \n", payR2)
	}

	t.Log(payR)
	t.Log(payR2)

}
