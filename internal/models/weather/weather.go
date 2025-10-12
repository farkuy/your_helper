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
	GetWeatherInfo(location string) (WeatherLocation, int, error)
}

type Model struct {
	Weather WeatherInfo
}

func Init(w WeatherInfo) *Model {
	return &Model{Weather: w}
}

func (m *Model) WeatherLocationInfo(location string) (string, error) {
	if location == "" {
		return "Введите город", nil
	}

	weather, statusCode, err := m.Weather.GetWeatherInfo(location)
	if err != nil {
		return "", err
	}

	//TODO: будумать над мидл вареной которая будет обрабатывать эти запросы
	switch statusCode {
	case 400:
		return "Данные по городу не найдены", nil
	case 500:
		return "Ошибка сервера, попробуйте позже", nil
	}

	return fmt.Sprintf("В городе %s по %s сейчас температура %.2f°C, ощущается как %.2f°C, влажность %.2f%%. Ветер дует со скоростью %.2f км/ч.",
		weather.Location.City,
		weather.Location.Region,
		weather.Current.TempC,
		weather.Current.FeelslikeC,
		weather.Current.Humidity,
		weather.Current.WindKph,
	), nil
}
