package go_weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kashifkhan0771/go-weather/config"
	"github.com/kashifkhan0771/go-weather/models"
)

// GetCurrentWeather return current weather response based on the option query
func (c WeatherAPIConfig) GetCurrentWeather(options Options) (*models.WeatherResponse, error) {
	if !validateQuery(options.Query) {
		return nil, fmt.Errorf("invalid query parameter")
	}

	url := fmt.Sprintf("%s%s?q=%s", config.BaseURL, config.CurrentWeatherJSON, options.Query)

	resp, err := c.makeRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Decode JSON response body
	var weatherResp models.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return nil, err
	}

	return &weatherResp, nil
}

func (c WeatherAPIConfig) WeatherHistory(options Options) (*models.WeatherResponse, error) {
	if !validateQuery(options.Query) {
		return nil, fmt.Errorf("invalid query parameter")
	} else if !validateDate(options.Date) {
		return nil, fmt.Errorf("invalid date parameter")
	}

	url := fmt.Sprintf("%s%s?q=%s&dt=%s", config.BaseURL, config.HistoryWeatherJSON, options.Query, options.Date.Format("2006-01-02"))

	resp, err := c.makeRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Decode JSON response body
	var weatherResp models.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return nil, err
	}

	return &weatherResp, nil
}

// makeRequest is a helper function for making API requests.
func (c WeatherAPIConfig) makeRequest(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("key", c.XApiKey)

	// Send request
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp, nil
}
