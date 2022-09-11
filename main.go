package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://api.weather.gov/gridpoints/RAH/74,58/forecast"

func main() {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", baseUrl, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	weather := WeatherData{}

	if body, err := io.ReadAll(resp.Body); err == nil {
		json.Unmarshal([]byte(body), &weather)

		fmt.Printf("Forecast last updated at %v\n\n", weather.Properties.Updated)

		for _, v := range weather.Properties.Periods {
			fmt.Printf("%v Temp: %v%v -Forecast: %v\n\n",
				v.Name,
				v.Temperature,
				v.TemperatureUnit,
				v.ShortForecast)
		}

	}
}
