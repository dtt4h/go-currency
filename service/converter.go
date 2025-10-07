package service

import (
	"fmt"

	"go-currency/client"
)

func ConvertCurrency(from, to string, amount float64) (float64, error) {
	rates, err := client.GetExchangeRates(from)
	if err != nil {
		return 0, fmt.Errorf("error while getting rates: %v", err)
	}

	fmt.Printf("Available rates: %+v\n", rates.Rates) // Отладка
	fmt.Printf("Looking for currency: %s\n", to)      // Отладка

	toRate, exists := rates.Rates[to]
	if !exists {
		return 0, fmt.Errorf("target currency '%s' not found in rates. Available: %v", to, getAvailableCurrencies(rates.Rates))
	}

	result := amount * toRate
	return result, nil
}

// Вспомогательная функция для отображения доступных валют
func getAvailableCurrencies(rates map[string]float64) []string {
	currencies := make([]string, 0, len(rates))
	for currency := range rates {
		currencies = append(currencies, currency)
	}
	return currencies
}
