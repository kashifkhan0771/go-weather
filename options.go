package go_weather

/*
Options represent the query parameters of the APIs.

For more details about each query parameters, please visit: https://www.weatherapi.com/docs/
*/
type Options struct {
	// Query parameter based on which data is sent back (Required)
	Query string `json:"q"`
}
