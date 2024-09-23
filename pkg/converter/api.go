package converter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type DataResponse struct {
	Data map[string]float64 `json:"data"`
}

const BASE_URL = "https://api.freecurrencyapi.com/v1/latest?apikey="

func FechRates(from string, to string) (map[string]float64, error) {

	apiKey := os.Getenv("FREE_CURRENCY_APP_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("Not API KEY found")
	}

	url := BASE_URL + apiKey + "&currencies=" + to + "&base_currency=" + from

	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch api rates : %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Api request failed with status: %v", response.StatusCode)
	}

	var dataResponse DataResponse

	if err := json.NewDecoder(response.Body).Decode(&dataResponse); err != nil {
		return nil, fmt.Errorf("Failed to decode API: %v", err)
	}

	return dataResponse.Data, nil
}

func ConvertValue(value float64, rate float64) float64 {
	return value * rate
}
