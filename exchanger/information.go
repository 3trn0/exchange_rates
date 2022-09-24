package exchanger

import (
	"fmt"
)

// PrintCurrenciesCodes function prints currencies short codes.
// This function helps user to understand that, for example, KWD stands for Kuwaiti Dinar.
func PrintCurrenciesCodes() error {
	inCodes, err := GetCurrenciesCodes()

	if err != nil {
		return err
	}

	fmt.Printf("Currencies Codes:\n")
	for code, name := range inCodes {
		fmt.Printf("%s: %s\n", code, name)
	}

	return nil
}

// PrintCurrentRelatedRates function prints currency related exchange rates.
// It takes curName as an argument, which should be provided as currency short code.
// For example: PrintCurrentRelatedRates("USD") will print exchange rates for United States Dollar.
func PrintCurrentRelatedRates(curName string) error {

	inRates, err := GetCurrentRelatedRates(curName)

	if err != nil {
		return err
	}

	fmt.Printf("Exchange Rates for %s currency:\n", curName)
	for code, value := range inRates {
		fmt.Printf("%s: %.10f\n", code, value)
	}

	return nil
}

// PrintMainRelatedRates function prints main currency related exchange rates.
// And to be more precise, only exchange rates related to UAH, USD, EUR, GBP, JPY, CNY.
func PrintMainRelatedRates(curName string) error {

	curRates, err := GetCurrentRelatedRates(curName)

	if err != nil {
		return err
	}

	inRates := ChooseMainRelatedRates(curRates)

	fmt.Printf("Exchange Rates for %s currency:\n", curName)
	for code, value := range inRates {
		fmt.Printf("%s: %.10f\n", code, value)
	}

	return nil
}
