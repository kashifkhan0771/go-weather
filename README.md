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
        config := weatherClient.WeatherAPIConfig{
            XApiKey: "<your_x_api_key>",
        }

        weather, err := config.GetCurrentWeather(client.Options{Query: "Peshawar"})
        if err != nil {
            panic(err)
        }

        fmt.Println(weather)
    }
</code>