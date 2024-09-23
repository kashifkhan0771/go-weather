# Go API Client

### API Docs:
[WeatherAPI Documentation](https://www.weatherapi.com/docs/)

### Usage:
### Download:
```bash
go get "github.com/kashifkhan0771/go-weather"
```
<code>

    package main

    import (
        "fmt"
        weatherClient "github.com/kashifkhan0771/go-weather"
    )

    func main() {
        config, err := weatherClient.NewWeatherAPIConfig(<Your_API_Key>)
        if err != nil {
            panic(err)
        }

        weather, err := config.GetCurrentWeather(weatherClient.Options{Query: "Paris"})
        if err != nil {
            panic(err)
        }

        fmt.Println(weather)
    }
</code>