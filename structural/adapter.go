// inspired by https://blog.matthiasbruns.com/golang-adapter-pattern from Matthias Bruns

// paypal.go
package structural

type PayPal struct{}

func (p *PayPal) MakePayment(amount float32) bool {
	// connect to PayPal and process payment
	return true
}

// amazon.go

type Amazon struct{}

func (a *Amazon) PayAmazon(amount float32) bool {
	// connect to Amazon and process payment
	return true
}

// gateway.go

type PaymentGateway interface {
	ProcessPayment(amount float32) bool
}

type PayPalAdapter struct {
	PayPal *PayPal
}

func (p *PayPalAdapter) ProcessPayment(amount float32) bool {
	return p.PayPal.MakePayment(amount)
}

type AmazonAdapter struct {
	Amazon *Amazon
}

func (a *AmazonAdapter) ProcessPayment(amount float32) bool {
	return a.Amazon.PayAmazon(amount)
}
