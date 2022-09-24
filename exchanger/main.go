package exchanger

import (
	"fmt"
)

// ConvertCurrencies function converts value 'v' from currency with short
// code 'codeFrom' to currency with short code 'codeTo'.
func ConvertCurrencies(codeFrom string, v float64, codeTo string) (float64, error) {
	curRate, err := GetCurrentRelatedRates(codeFrom)
	if err != nil {
		return -1, err
	}

	rate, ok := curRate[codeTo]
	if !ok {
		return 0, fmt.Errorf("there is no available currency with %s short code", codeTo)
	}

	cur := NewCurrency(v)

	cur.Multiply(rate).Float64()

	fmt.Printf("%.2f %s ===> %s %s\n", v, codeFrom, cur.String(), codeTo)

	return cur.Float64(), nil

}
