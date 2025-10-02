package client

import (
	"encoding/json"
	"fmt"
	"go-currency/models"
	"io"
	"net/http"
)

const baseURL = "https://api.exchangerate-api.com/v4/latest/"

func GetExchangeRates(baseCurrency string) (*models.ExchangeRates, error) {
	url := baseURL + baseCurrency
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error while request to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error API: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading  response body: %v", err)
	}

	var rates models.ExchangeRates

	err = json.Unmarshal(body, &rates)
	if err != nil {
		return nil, fmt.Errorf("error while parsing JSON: %v", err)
	}
	return &rates, nil
}
