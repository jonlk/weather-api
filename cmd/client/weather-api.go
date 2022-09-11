package main

import "github/jonlk/weather-api/pkg/weather"

func main() {
	office := "RAH"
	gridX := 74
	gridY := 58

	weather.GenerateForecast(&office, &gridX, &gridY)
}
