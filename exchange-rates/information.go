package exchange_rates

import (
	"fmt"
	"log"
)

//PrintCurrenciesCodes function prints currencies short codes.
//This function helps user to understand that, for example, KWD stands for Kuwaiti Dinar.
func PrintCurrenciesCodes() {
	inCodes, err := GetCurrenciesCodes()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Currencies Codes:\n")
	for code, name := range inCodes {
		fmt.Printf("%s: %s\n", code, name)
	}
}

//PrintCurrentRelatedRates function prints currency related exchange rates.
//It takes curName as an argument, which should be provided as currency short code.
//For example: PrintCurrentRelatedRates("USD") will print exchange rates for United States Dollar.
func PrintCurrentRelatedRates(curName string) {

	inRates, err := GetCurrentRelatedRates(curName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Exchange Rates for %s currency:\n", curName)
	for code, value := range inRates {
		fmt.Printf("%s: %.10f\n", code, value)
	}
}

//PrintMainRelatedRates function prints main currency related exchange rates.
//And to be more precise, only exchange rates related to UAH, USD, EUR, GBP, JPY, CNY.
func PrintMainRelatedRates(curName string) {

	curRates, err := GetCurrentRelatedRates(curName)

	if err != nil {
		log.Fatal(err)
	}

	inRates := ChooseMainRelatedRates(curRates)

	fmt.Printf("Exchange Rates for %s currency:\n", curName)
	for code, value := range inRates {
		fmt.Printf("%s: %.10f\n", code, value)
	}
}
