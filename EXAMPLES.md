


# Examples for Go-API Client

## Example 1: Fetching Current Weather
```go
package main

import (
    "fmt"
    "github.com/kashifkhan0771/go-weather"
)

func main() {
    client := weather.NewClient("YOUR_API_KEY")
    currentWeather, err := client.GetCurrentWeather("New York")
    if err != nil {
        fmt.Println("Error fetching weather:", err)
        return
    }

    fmt.Printf("Current temperature in New York: %f°C\n", currentWeather.Temperature)
}

```

## Example 2: Fetching a 5-Day Forecast
```go
Copy code
package main

import (
    "fmt"
    "github.com/kashifkhan0771/go-weather"
)

func main() {
    client := weather.NewClient("YOUR_API_KEY")
    forecast, err := client.GetForecast("Tokyo", 5)
    if err != nil {
        fmt.Println("Error fetching forecast:", err)
        return
    }

    for _, day := range forecast.Days {
        fmt.Printf("Date: %s, Temp: %f°C\n", day.Date, day.Temperature)
    }
}
```

## Example 3: Handling Errors
```go
Copy code
package main

import (
    "fmt"
    "github.com/kashifkhan0771/go-weather"
)

func main() {
    client := weather.NewClient("INVALID_API_KEY")
    _, err := client.GetCurrentWeather("Paris")
    if err != nil {
        fmt.Println("Error:", err)
    }

}
```
