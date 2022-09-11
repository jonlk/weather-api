package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GenerateForecast(office *string, gridX *int, gridY *int) {

	baseUrl := "https://api.weather.gov/gridpoints/" +
		*office + "/" +
		fmt.Sprint(*gridX) + "," +
		fmt.Sprint(*gridY) +
		"/forecast"

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
