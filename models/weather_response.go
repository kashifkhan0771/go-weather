package models

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
