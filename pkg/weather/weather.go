package weather

type Forecast struct {
	Properties Properties
}

type Properties struct {
	Updated           string
	Units             string
	ForecastGenerator string
	GeneratedAt       string
	UpdateTime        string
	ValidTimes        string
	Elevation         Elevation
	Periods           []Period
}

type Elevation struct {
	UnitCode string
	Value    float64
}

type Period struct {
	Number           int
	Name             string
	StartTime        string
	EndTime          string
	IsDaytime        string
	Temperature      int
	TemperatureUnit  string
	TemperatureTrend string
	WindSpeed        string
	WindDirection    string
	Icon             string
	ShortForecast    string
	DetailedForecast string
}
