package messages

import (
	"strings"
	"your_helper/internal/models/weather"
)

type Model struct {
	WeatherModel weather.Model
}

func Init(w weather.Model) *Model {
	return &Model{WeatherModel: w}
}

func (m *Model) CreateAnswer(text string) (string, error) {
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
		default:
			res, err = "Проверьте вводимую команду", nil
		}
	}

	if len(words) > 1 {
		switch words[0] {
		case "/weather":
			res, err = m.WeatherModel.WeatherLocationInfo(words[1:])
		default:
			res, err = "Проверьте вводимую команду", nil
		}
	}

	return res, err
}
