package exchange_rates

import (
	"errors"
	"fmt"
	"strconv"
) /*exchange_rates*/

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//GetJSON function is a base function to make requests to https://api.coinbase.com API.
func GetJSON(url string, method string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(res.Body)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//GetCurrenciesCodes function returns currencies short codes.
func GetCurrenciesCodes() (map[string]string, error) {
	baseUrl := "https://api.coinbase.com/v2/currencies"

	body, err := GetJSON(baseUrl, "GET")

	if err != nil {
		return nil, err
	}

	var response map[string][]map[string]string

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	resCodes := make(map[string]string)

	for _, cur := range response["data"] {
		resCodes[cur["id"]] = cur["name"]
	}

	return resCodes, nil
}

//GetCurrentRelatedRates function returns current exchange rates related to currency,
//which short code is provided.
func GetCurrentRelatedRates(curName string) (map[string]float64, error) {

	codes, err := GetCurrenciesCodes()

	if err != nil {
		return nil, err
	}

	if _, ok := codes[curName]; !ok {
		return nil, errors.New(fmt.Sprintf("there is no available currency with %s short code", curName))
	}

	baseUrl := "https://api.coinbase.com/v2/exchange-rates?currency=" + strings.ToUpper(curName)

	body, err := GetJSON(baseUrl, "GET")

	if err != nil {
		return nil, err
	}

	var response map[string]map[string]interface{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	curMap := response["data"]["rates"].(map[string]interface{})

	curFloatMap, err := convRatesToFloat(curMap)

	resMap := make(map[string]float64)

	for code := range codes {
		resMap[code] = curFloatMap[code]
	}

	if err != nil {
		return nil, err
	}

	return resMap, nil
}

//ChooseMainRelatedRates function chooses and returns main exchange rates
//that are related to a provided currency short code.
//And to be more precise, by main we mean only rates related to UAH, USD, EUR, GBP, JPY, CNY.
func ChooseMainRelatedRates(inRates map[string]float64) map[string]float64 {
	var resMap = map[string]float64{"UAH": 0.0, "USD": 0.0, "EUR": 0.0, "GBP": 0.0, "JPY": 0.0, "CNY": 0.0}

	for cur := range resMap {

		resMap[cur] = inRates[cur]
	}

	return resMap
}

//convRatesToFloat function is the inner-used function that converts exchange rates to float64
//so that we can perform some actions with them, as it was simple float64 number.
func convRatesToFloat(inRates map[string]interface{}) (map[string]float64, error) {
	resMap := make(map[string]float64)
	for cur, strPrice := range inRates {

		price, err := strconv.ParseFloat(fmt.Sprint(strPrice), 64)
		if err != nil {
			return nil, err
		}

		resMap[cur] = price
	}

	return resMap, nil
}
