package location

import (
	"fmt"
	"log/slog"
)

type Location struct {
	Id       int
	Location string
}

type LocationInfo interface {
	GetLocationInfo(id int) (Location, error)
	PostLocationInfo(id int, location string) (Location, error)
}

type Model struct {
	Location LocationInfo
}

func Init(l LocationInfo) *Model {
	return &Model{Location: l}
}

// TODO: попроблвать побаловать с sql инъекциями
func (m *Model) AddLocation(id int, location string) string {
	locInfo, err := m.Location.PostLocationInfo(id, location)
	if err != nil {
		slog.Error(err.Error())
		return fmt.Sprintf("Произогла ошибка добавление города %s", location)
	}

	return fmt.Sprintf("Вы успешно добавили город %s своим дефолтным", locInfo.Location)
}

func (m *Model) GetLocation(id int) string {
	locInfo, err := m.Location.GetLocationInfo(id)
	if err != nil {
		slog.Error(err.Error())
		return fmt.Sprintln("Произошла ошибка при поиcке города в вашем профиле")
	}

	return fmt.Sprintf("Ваш город %s", locInfo.Location)
}
