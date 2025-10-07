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
		return nil, fmt.Errorf("error while request to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error API: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %v", err)
	}

	var apiResponse struct {
		Base  string             `json:"base"`
		Rates map[string]float64 `json:"rates"`
	}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("error while parsing JSON: %v", err)
	}

	return &models.ExchangeRates{
		Base:  apiResponse.Base,
		Rates: apiResponse.Rates,
	}, nil
}
