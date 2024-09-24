package go_weather

import (
	"errors"
	"net/http"
	"time"

	"github.com/kashifkhan0771/go-weather/config"
)

// WeatherAPIConfig has the configurations required to call the APIs
type WeatherAPIConfig struct {
	// XApiKey can be obtained from: https://www.weatherapi.com/ after you log in
	XApiKey    string       `json:"key"`
	HttpClient *http.Client `json:"-"`
}

// NewWeatherAPIConfig creates a WeatherAPIConfig with apiKey passed.
func NewWeatherAPIConfig(apiKey string) (*WeatherAPIConfig, error) {
	if apiKey == "" {
		return nil, errors.New("API key cannot be empty")
	}

	return &WeatherAPIConfig{
		XApiKey:    apiKey,
		HttpClient: &http.Client{Timeout: time.Second * config.APIsTimeout},
	}, nil
}
