package weather

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	weather_model "your_helper/internal/models/weather"
)

const (
	basePath          = "https://api.weatherapi.com/v1"
	getWeatherInfoKey = "/current.json"
)

type Transport struct {
	apiKey string
}

func Init(key string) *Transport {
	return &Transport{apiKey: key}
}

func (m *Transport) Get(path string, params map[string]string) (res *http.Response, err error) {
	req, err := http.NewRequest("GET", basePath+path, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	q.Add("lang", "ru")
	q.Add("key", m.apiKey)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req)

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (m *Transport) GetWeatherInfo(location string) (weather_model.WeatherLocation, int, error) {
	log := slog.With("PATH: transports/rest/weather")

	var weather weather_model.WeatherLocation
	localParams := map[string]string{
		"q": location,
	}

	res, err := m.Get(getWeatherInfoKey, localParams)
	statusCode := res.StatusCode
	if err != nil {
		log.Error("Ошибка при запросе: %v", err)
		return weather_model.WeatherLocation{}, statusCode, err
	}
	if res.StatusCode == 400 {
		log.Warn("Данные по локации не найдены")
		return weather_model.WeatherLocation{}, statusCode, err
	}

	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		log.Error("Ошибка парсинга json", location)
		return weather_model.WeatherLocation{}, statusCode, err
	}

	return weather, statusCode, err
}
