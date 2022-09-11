package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://api.weather.gov/gridpoints/RAH/74,58/forecast"

func GenerateForecast() {

	req, _ := http.NewRequest("GET", baseUrl, nil)

	client := &http.Client{}

	if resp, err := client.Do(req); err == nil {

		if body, err := io.ReadAll(resp.Body); err == nil {
			weather := WeatherData{}
			json.Unmarshal([]byte(body), &weather)
			printOutput(&weather)

		} else {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println(err.Error())
	}
}

func printOutput(weather *WeatherData) {

	fmt.Printf("\nForecast last updated at %v\n\n", weather.Properties.Updated)

	for _, v := range weather.Properties.Periods {
		fmt.Printf("%v Temp: %v%v -Forecast: %v\n",
			v.Name,
			v.Temperature,
			v.TemperatureUnit,
			v.ShortForecast)
	}
}