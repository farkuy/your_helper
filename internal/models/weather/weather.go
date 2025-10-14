package weather

import (
	"fmt"
	"log/slog"
	"sync"
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

func (m *Model) WeatherLocationInfo(locations []string) (string, error) {
	message := ""

	if len(locations) == 0 {
		return "Введите город", nil
	}
	if len(locations) > 30 {
		return "Максимум 30 городов за раз", nil
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	locationsCh := make(chan string, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(locationsCh)

		for _, val := range locations {
			locationsCh <- val
		}
	}()

	wg.Add(3)
	for range 3 {
		go func() {
			defer wg.Done()

			for val := range locationsCh {
				textUnfo := "\n" + getLocInfo(val, m)
				mu.Lock()
				message += textUnfo
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	return message, nil
}

func getLocInfo(location string, m *Model) string {
	weather, statusCode, err := m.Weather.GetWeatherInfo(location)
	if err != nil {
		slog.Error(err.Error())
		return fmt.Sprintf("Произогла ошибка по городу %s", location)
	}

	//TODO: будумать над мидл вареной которая будет обрабатывать эти запросы
	switch statusCode {
	case 400:
		return fmt.Sprintf("Данные по городу не найдены %s", location)
	case 500:
		return fmt.Sprintf("Ошибка сервера по городу %s, попробуйте позже", location)
	}

	return fmt.Sprintf("В городе %s по %s сейчас температура %.2f°C, ощущается как %.2f°C, влажность %.2f%%. Ветер дует со скоростью %.2f км/ч.",
		weather.Location.City,
		weather.Location.Region,
		weather.Current.TempC,
		weather.Current.FeelslikeC,
		weather.Current.Humidity,
		weather.Current.WindKph,
	)
}
