package main

import (
	"fmt"
	"github/jonlk/weather-api/pkg/weather"
	"os"
)

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
