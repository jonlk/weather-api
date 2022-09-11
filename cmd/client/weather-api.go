package main

import (
	"fmt"
	"github/jonlk/weather-api/pkg/weather"
	"os"
)

// TODO: Need to add an Api call and also look at other features
//
//	on the weather.gov api
func main() {

	if len(os.Args) != 2 || (os.Args[1] != "short" && os.Args[1] != "long") {
		fmt.Println("USAGE: 1 ARGUMENT REQUIRED - 'long' for detailed forecast, 'short' for abbreviated forecast")
		return
	}

	office := "RAH"
	gridX := 74
	gridY := 58

	f := weather.NewForecaster()
	f.GenerateForecast(&office, &gridX, &gridY, os.Args[1])
}
