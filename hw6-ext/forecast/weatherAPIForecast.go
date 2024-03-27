package forecast

type WApi struct {
}

func (r *WApi) GetTomorrowWeather() string {
	// https://api.open-meteo.com/v1/forecast?latitude=50.4547&longitude=30.5238&hourly=temperature_2m&forecast_days=1
	return "get json parse"
}
