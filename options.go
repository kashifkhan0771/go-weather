package go_weather

import "time"

/*
Options represent the query parameters of the APIs.

For more details about each query parameters, please visit: https://www.weatherapi.com/docs/
*/
type Options struct {
	// Query parameter based on which data is sent back (Required)
	Query string `json:"q"`
	// Date is required for history and future API, it restricts date output for Forecast and History API
	Date time.Time `json:"dt"`
}
