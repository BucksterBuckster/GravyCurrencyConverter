package gravycurrencyconverter

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

const availableCurrenciesURL string = "https://api.exchangeratesapi.io/latest?base=USD"

const convertCurrencyURL string = "https://api.exchangeratesapi.io/latest?base=USD&symbols=%s,%s"

// Currency contains ID and Description of an Currency.
type Currency struct {
	Base  string
	Date  string
	Rates map[string]float64
}

// NewCurrency Create an instance of Currency.
//func NewCurrency(currency string) Currency {
//	return Currency{ID: currency}
//}

// AvailableCurrencies returns an array with all currencies that can be used.
func AvailableCurrencies() (Currency, error) {
	resp, err := newRequest().Get(availableCurrenciesURL)
	if err != nil {
		fmt.Println("There was an error with the request: ", err)
	}

	curList := Currency{}
	err = json.Unmarshal(resp, &curList)
	if err != nil {
		fmt.Println("There was an error with the JSON marshalling: ", err)
	}

	return curList, nil
}

type convertCurrencyResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

// ConvertCurrency Converts an amount (Decimal) from one currency to another currency.
func ConvertCurrency(from, to string, amount decimal.Decimal) (decimal.Decimal, error) {
	endpoint := fmt.Sprintf(convertCurrencyURL, from, to)
	resp, err := newRequest().Get(endpoint)
	if err != nil {
		return decimal.NewFromFloat(0), err
	}

	ccResp := convertCurrencyResponse{}
	err = json.Unmarshal(resp, &ccResp)
	if err != nil {
		return decimal.NewFromFloat(0), err
	}

	amountToFloat, _ := amount.Float64()

	convertedAmount := ccResp.Rates[to] * amountToFloat
	//if ccResp.Error != "" {
	//	return decimal.NewFromFloat(0), errors.New(ccResp.Error)
	//}
	return decimal.NewFromFloat(convertedAmount), nil
}
