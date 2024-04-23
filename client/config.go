package client

// WeatherAPIConfig has the configurations required to call the APIs
type WeatherAPIConfig struct {
	// XApiKey can be obtained from: https://www.weatherapi.com/ after you log in
	XApiKey string `json:"key"`
}
