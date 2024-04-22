package models

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMPH          float64   `json:"wind_mph"`
	WindKPH          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMB       float64   `json:"pressure_mb"`
	PressureIN       float64   `json:"pressure_in"`
	PrecipMM         float64   `json:"precip_mm"`
	PrecipIN         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelsLikeC       float64   `json:"feelslike_c"`
	FeelsLikeF       float64   `json:"feelslike_f"`
	VisibilityKM     float64   `json:"vis_km"`
	VisibilityMiles  float64   `json:"vis_miles"`
	UVIndex          float64   `json:"uv"`
	GustMPH          float64   `json:"gust_mph"`
	GustKPH          float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}
