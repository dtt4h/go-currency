package service

import (
	"fmt"
	"go-currency/models"
	"go-currency/client"
)

//Converts an amount from one currency to another
func ConvertCurrency(from, to string, amount float64) (float64, error) {
	rates, err := client.GetExchangeRates(from)
	if err != nil {
		return 0, fmt.Errof("error while getting rates: %v", err)
	}

	fromRate, exists := rates.Rates[from]

	if !exists {
		return 0, fmt.Errof("exchange %s not found in rates", from)
	}

	toRate, exists := rates.Rates[to]

	if !exists {
		return 0, fmt.Errof("exchange %s not found in rates", to)
	}

	result := amount * (toRate / fromRate)
	return result, nil
}