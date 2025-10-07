package client

import (
	"encoding/json"
	"fmt"
	"go-currency/models"
	"io"
	"net/http"
)

const baseURL = "https://api.frankfurter.app/latest"

func GetExchangeRates(baseCurrency string) (*models.ExchangeRates, error) {
	url := fmt.Sprintf("%s?from=%s", baseURL, baseCurrency)

	resp, err := http.Get(url)
	if err != nil {
		// Fallback to mock data if API fails
		return getMockRates(baseCurrency), nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Fallback to mock data if API fails
		return getMockRates(baseCurrency), nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return getMockRates(baseCurrency), nil
	}

	var apiResponse struct {
		Base  string             `json:"base"`
		Rates map[string]float64 `json:"rates"`
	}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return getMockRates(baseCurrency), nil
	}

	return &models.ExchangeRates{
		Base:  apiResponse.Base,
		Rates: apiResponse.Rates,
	}, nil
}

func getMockRates(baseCurrency string) *models.ExchangeRates {
	rates := map[string]float64{
		"USD": 1.0,
		"EUR": 0.92,
		"GBP": 0.79,
		"JPY": 148.50,
		"CAD": 1.35,
		"RUB": 90.0,
	}
	return &models.ExchangeRates{
		Base:  baseCurrency,
		Rates: rates,
	}
}
