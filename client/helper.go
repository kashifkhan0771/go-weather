package client

import "net/http"

func (c *WeatherAPIConfig) makeRequest(method, url string) (*http.Response, error) {
	client := &http.Client{}

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
