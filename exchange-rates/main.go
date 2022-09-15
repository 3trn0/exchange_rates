package exchange_rates

import (
	"errors"
	"fmt"
)

//ConvertCurrencies function converts value 'v' from currency with short
//code 'code1' to currency with short code 'code2'.
func ConvertCurrencies(code1 string, v float64, code2 string) (float64, error) {
	curRate, err := GetCurrentRelatedRates(code1)
	if err != nil {
		return -1, err
	}

	rate, ok := curRate[code2]
	if !ok {
		return -1, errors.New(fmt.Sprintf("there is no available currency with %s short code", code2))
	}

	cur := ToCurrency(v)

	cur.Multiply(rate).ToFloat64()

	fmt.Printf("%.2f %s ===> %s %s\n", v, code1, cur.ToString(), code2)

	return cur.ToFloat64(), nil

}
