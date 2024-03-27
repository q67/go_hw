package main

type tomorrowWeather interface {
	GetTomorrowWeather() string
}

func printForecast(t tomorrowWeather) string {
	return t.GetTomorrowWeather()
}

func main() {
	f :=
}
