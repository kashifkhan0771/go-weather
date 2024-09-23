package config

const (
	// BaseURL - is the base url of the all APIs
	BaseURL = "https://api.weatherapi.com/v1"

	// API EndPoints
	CurrentWeatherJSON = "/current.json"
	HistoryWeatherJSON = "/history.json"

	// APIsTimeout is a timeout for all the API request to api.weatherapi.com
	APIsTimeout = 30
)
