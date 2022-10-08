package exchanger

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// getJSON function is a base private function to make requests to https://api.coinbase.com API.
func getJSON(url string, method string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf(http.StatusText(res.StatusCode))
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetCurrenciesCodes function returns currencies short codes.
func GetCurrenciesCodes() (map[string]string, error) {
	baseUrl := "https://api.coinbase.com/v2/currencies"

	body, err := getJSON(baseUrl, http.MethodGet)

	if err != nil {
		return nil, err
	}

	response := CurrenciesCodesResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	resCodes := make(map[string]string)

	for _, cur := range response.Data {
		resCodes[cur.Id] = cur.Name
	}

	return resCodes, nil
}

// GetCurrentRelatedRates function returns current exchange rates related to currency,
// which short code is provided.
func GetCurrentRelatedRates(curName string) (map[string]float64, error) {

	codes, err := GetCurrenciesCodes()

	if err != nil {
		return nil, err
	}

	if _, ok := codes[curName]; !ok {
		return nil, fmt.Errorf("there is no available currency with %s short code", curName)
	}

	baseUrl := "https://api.coinbase.com/v2/exchange-rates?currency=" + strings.ToUpper(curName)

	body, err := getJSON(baseUrl, http.MethodGet)

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

// ChooseMainRelatedRates function chooses and returns main exchange rates
// that are related to a provided currency short code.
// And to be more precise, by main we mean only rates related to UAH, USD, EUR, GBP, JPY, CNY.
func ChooseMainRelatedRates(inRates map[string]float64) map[string]float64 {
	var resMap = map[string]float64{"UAH": 0.0, "USD": 0.0, "EUR": 0.0, "GBP": 0.0, "JPY": 0.0, "CNY": 0.0}

	for cur := range resMap {

		resMap[cur] = inRates[cur]
	}

	return resMap
}

// convRatesToFloat function is the inner-used function that converts exchange rates to float64
// so that we can perform some actions with them, as it was simple float64 number.
func convRatesToFloat(inRates map[string]interface{}) (map[string]float64, error) {
	resMap := make(map[string]float64)
	for cur, strPrice := range inRates {
		str, ok := strPrice.(string)
		if !ok {
			return nil, fmt.Errorf("can not convert %s currency", cur)
		}
		price, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}

		resMap[cur] = price
	}

	return resMap, nil
}
