package weather

import (
	"fmt"
)

type Printer struct{}

func newPrinter() *Printer {
	return &Printer{}
}

func (p *Printer) printShortOutput(weather *WeatherData) {

	fmt.Printf("\nForecast last updated at %v\n\n",
		weather.Properties.Updated)

	for _, v := range weather.Properties.Periods {
		fmt.Printf("%v Temp: %v%v -Forecast: %v\n\n",
			v.Name,
			v.Temperature,
			v.TemperatureUnit,
			v.ShortForecast)
	}
}

func (p *Printer) printLongOutput(weather *WeatherData) {

	fmt.Printf("\nForecast last updated at %v\n\n",
		weather.Properties.Updated)

	for _, v := range weather.Properties.Periods {
		fmt.Printf("%v -Forecast: %v\n\n",
			v.Name,
			v.DetailedForecast)
	}
}
