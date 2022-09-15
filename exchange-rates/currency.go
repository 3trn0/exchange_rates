package exchange_rates

import "fmt"

// Currency type represents any currency value.
//It is used to шьздуьуте more accurate calculations on currencies.
type Currency int64

//ToCurrency function converts float64 value to Currency value.
func ToCurrency(v float64) Currency {
	x := v*100 + 0.5
	return Currency(x)
}

//ToFloat64 method returns float64 Currency representation.
func (c *Currency) ToFloat64() float64 {
	return float64(*c) / 100
}

//ToString method returns string Currency representation.
func (c *Currency) ToString() string {
	x := float64(*c) / 100
	return fmt.Sprintf("%.2f", x)
}

//Multiply method implements Currency multiplication by some float64 value.
//It saves changes to Currency value.
func (c *Currency) Multiply(v float64) *Currency {
	x := float64(*c)*v + 0.5
	*c = Currency(x)
	return c
}
