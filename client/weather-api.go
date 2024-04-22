package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kashifkhan0771/go-weather/config"
	"github.com/kashifkhan0771/go-weather/models"
)

type WeatherAPIConfig struct {
	XApiKey string `json:"key"`
}

func (c *WeatherAPIConfig) GetCurrentWeather(query string) (*models.WeatherResponse, error) {
	if query == "" {
		return nil, fmt.Errorf("query is empty")
	}

	url := fmt.Sprintf("%s%s?q=%s", config.BaseURL, config.CurrentWeatherJSON, query)

	// Create HTTP client
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("key", c.XApiKey)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
		return nil, err
	}

	// Decode JSON response body
	var weatherResp *models.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return nil, err
	}

	return weatherResp, nil
}
