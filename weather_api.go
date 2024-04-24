package go_weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kashifkhan0771/go-weather/config"
	"github.com/kashifkhan0771/go-weather/models"
)

// GetCurrentWeather return current weather response based on the option query
func (c WeatherAPIConfig) GetCurrentWeather(options Options) (*models.WeatherResponse, error) {
	if options.Query == "" {
		return nil, fmt.Errorf("query is empty")
	}

	url := fmt.Sprintf("%s%s?q=%s", config.BaseURL, config.CurrentWeatherJSON, options.Query)

	resp, err := c.makeRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
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

// makeRequest is a helper function for making API requests.
func (c WeatherAPIConfig) makeRequest(method, url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * config.APIsTimeout,
	}

	req, err := http.NewRequest(method, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("key", c.XApiKey)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
