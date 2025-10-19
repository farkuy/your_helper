package messages

import (
	"strings"
	"your_helper/internal/models/location"
	"your_helper/internal/models/weather"
)

type Model struct {
	WeatherModel  weather.Model
	LocationModel location.Model
}

func Init(w *weather.Model, l *location.Model) *Model {
	return &Model{WeatherModel: *w, LocationModel: *l}
}

func (m *Model) CreateAnswer(userId int64, text string) (string, error) {
	var (
		res string = ""
		err error  = nil
	)

	words := strings.Split(text, " ")

	if len(words) == 0 {
		res, err = "Введите команду", nil
	}

	if len(words) == 1 {
		switch words[0] {
		case "/hi":
			res, err = "Жми екарный бабай", nil
		case "/getLocation":
			res, err = m.LocationModel.GetLocation(userId)
		default:
			res, err = "Проверьте вводимую команду", nil
		}
	}

	if len(words) > 1 {
		switch words[0] {
		case "/weather":
			res, err = m.WeatherModel.WeatherLocationInfo(words[1:])
		case "/addLocation":
			res, err = m.LocationModel.AddLocation(userId, words[1])
		default:
			res, err = "Проверьте вводимую команду", nil
		}
	}

	return res, err
}
