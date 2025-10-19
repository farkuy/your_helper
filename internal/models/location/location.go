package location

import (
	"fmt"
	"log/slog"
)

type Location struct {
	Id       int    `db:"id" json:"id"`
	Location string `db:"location" json:"location"`
}

type LocationInfo interface {
	PostLocationInfo(id int64, location string) (Location, error)
	GetLocationInfo(id int64) (Location, error)
}

type Model struct {
	Location LocationInfo
}

func Init(l LocationInfo) *Model {
	return &Model{Location: l}
}

// TODO: попроблвать побаловать с sql инъекциями
func (m *Model) AddLocation(id int64, location string) (string, error) {
	if location == "" {
		return fmt.Sprintln("Введи город в запросе после команды", location), nil
	}

	fmt.Println(id, location, &m, m)
	locInfo, err := m.Location.PostLocationInfo(id, location)
	if err != nil {
		slog.Error(err.Error())
		return fmt.Sprintf("Произогла ошибка добавление города %s", location), err
	}

	return fmt.Sprintf("Вы успешно добавили город %s своим дефолтным", locInfo.Location), nil
}

func (m *Model) GetLocation(id int64) (string, error) {
	locInfo, err := m.Location.GetLocationInfo(id)
	if err != nil {
		slog.Error(err.Error())
		return fmt.Sprintln("Произошла ошибка при поиcке города в вашем профиле"), err
	}

	return fmt.Sprintf("Ваш город %s", locInfo.Location), nil
}
