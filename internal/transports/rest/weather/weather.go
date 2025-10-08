package weather

import "net/http"

const (
	getWeatherInfoKey = "/api.weatherapi.com/v1/current.json?q="
)

type Model struct {
	apiKey string
}

func Init(key string) Model {
	return Model{apiKey: key}
}

func (m *Model) AddEndKey() string {
	return "&key=" + m.apiKey
}

// TODO: добавить возможность менять язык
func AddEndLang() string {
	return "&lang=ru"
}

func (m *Model) Get(path string) (resp *http.Response, err error) {
	resp, err = http.Get("https://" + path + AddEndLang() + m.AddEndKey())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, err
}

func (m *Model) GetWeatherInfo(location string) (string, error) {
	resp, err := m.Get(getWeatherInfoKey + location)

}
