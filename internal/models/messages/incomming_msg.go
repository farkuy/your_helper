package messages

import "your_helper/internal/transports/rest/weather"

type Model struct{}

func Init() *Model {
	return &Model{}
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
		res, err = weather.GetWeatherInfo("Дзержинск")
	default:
		res, err = "Я хз что это", nil
	}

	return res, err
}
