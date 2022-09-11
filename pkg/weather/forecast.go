package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Forecaster struct{}

func NewForecaster() *Forecaster {
	return &Forecaster{}
}

func (f *Forecaster) GenerateForecast(office *string, gridX *int, gridY *int, outputType string) {

	baseUrl := "https://api.weather.gov/gridpoints/" +
		*office + "/" +
		fmt.Sprint(*gridX) + "," +
		fmt.Sprint(*gridY) +
		"/forecast"

	req, _ := http.NewRequest("GET", baseUrl, nil)

	client := &http.Client{}

	if resp, err := client.Do(req); err == nil {

		if body, err := io.ReadAll(resp.Body); err == nil {

			weather := Forecast{}

			json.Unmarshal([]byte(body), &weather)

			p := newPrinter()

			switch outputType {
			case "short":
				p.printShortOutput(&weather)
			case "long":
				p.printLongOutput(&weather)
			default:
				p.printShortOutput(&weather)
			}

		} else {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println(err.Error())
	}
}
