# Go API Weather

[![Go Version](https://img.shields.io/github/go-mod/go-version/kashifkhan0771/go-weather)](https://golang.org/)
[![License](https://img.shields.io/github/license/kashifkhan0771/go-weather)](LICENSE)

Go-Weather is a simple and efficient Go library for fetching weather information from  https://www.weatherapi.com . This library allows developers to seamlessly integrate weather data into their applications.

## Features
- Fetch current weather for any location.
- Retrieve forecast data for up to 7 days.
- Supports WeatherAPI integration for reliable weather data.
- Lightweight

## Installation
In order to use Go-Weather, you need to have Go installed on your system. The library requires Go version 1.22.0 or higher.

```shell
go get github.com/kashifkhan0771/go-weather
```

## Usage
Import the library in your Go application:


```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/go-weather"
)

func main() {
    client := weather.NewClient("YOUR_API_KEY")
    currentWeather, err := client.GetCurrentWeather("London")
    if err != nil {
        fmt.Println("Error fetching weather:", err)
        return
    }

    fmt.Printf("Current temperature in London: %f°C\n", currentWeather.Temperature)
}
```

## Configuration
Set up your client with the API key provided by the weather service:

```
client := weather.NewClient("YOUR_API_KEY")
```

## Supported Functions
- `GetCurrentWeather(location string)`: Fetch current weather for a given location.
- `GetForecast(location string, days int)`: Retrieve weather forecast for a specified number of days.

## Examples
Detailed usage examples for each function are provided in the [EXAMPLES.md](EXAMPLES.md) file.


## Contributing
Contributions are welcome! Please follow the steps below:

- Fork the repository.
- Create a feature branch: git checkout -b feature-name.
- Commit your changes: git commit -m 'Add new feature'.
- Push to your branch: git push origin feature-name.
- Open a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/kashifkhan0771/go-weather/blob/main/LICENSE) file for details.

## Acknowledgments
Thanks to WeatherAPI for data integration.
