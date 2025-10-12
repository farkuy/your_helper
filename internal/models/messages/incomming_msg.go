package messages

import "your_helper/internal/models/weather"

type Model struct {
	weather weather.Model
}

func Init(w weather.Model) *Model {
	return &Model{weather: w}
}

func (m *Model) CreateAnswer(text string) (string, error) {
	var (
		res string = ""
		err error  = nil
	)

	switch text {
	case "/hi":
		res, err = "Жми екарный бабай", nil
	case "/weather":
		res, err = m.weather.WeatherLocationInfo("Дзержинск")
	default:
		res, err = "Я хз что это", nil
	}

	return res, err
}
