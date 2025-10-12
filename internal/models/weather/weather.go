package weather

import (
	"fmt"
)

type LocationInfo struct {
	City    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Time    string `json:"localtime"`
}

type Current struct {
	TempC      float32 `json:"temp_c"`
	FeelslikeC float32 `json:"feelslike_c"` //Ощущается как
	Humidity   float32 `json:"humidity"`    //Влажность
	WindKph    float32 `json:"wind_kph"`    //Скорость ветра км/ч
}

type WeatherLocation struct {
	Current  Current      `json:"current"`
	Location LocationInfo `json:"location"`
}

type WeatherInfo interface {
	GetWeatherInfo(location string) (WeatherLocation, error)
}

type Model struct {
	Weather WeatherInfo
}

func Init(w WeatherInfo) *Model {
	return &Model{Weather: w}
}

func (m *Model) WeatherLocationInfo(location string) (string, error) {
	weather, err := m.Weather.GetWeatherInfo(location)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("В городе %s сейчас температура %d°C, ощущается как %d°C, влажность %d%%. Погода %s, ветер дует со скоростью %d км/ч.",
		weather.Location.City,
		weather.Current.TempC,
		weather.Current.FeelslikeC,
		weather.Current.Humidity,
		weather.Current.WindKph,
	), nil
}
