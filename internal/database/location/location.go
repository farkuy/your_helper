package location_bd

import (
	"context"
	"errors"
	"log/slog"
	"your_helper/internal/models/location"

	"github.com/jackc/pgx/v4"
)

const (
	addLocation    = "INSERT INTO users (id, location) VALUES ($1, $2) RETURNING id, location"
	updateLocation = "UPDATE users SET location = $1 WHERE id = $2 RETURNING id, location"
	checkUserInfo  = "SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)"
	getUserInfo    = "SELECT * FROM users WHERE id = $1"
)

type Transport struct {
	bd *pgx.Conn
}

func Init(bd *pgx.Conn) Transport {
	return Transport{bd: bd}
}

var log = slog.With("PATH: internal/database/location/location")

func (t *Transport) PostLocationInfo(id int64, newLocation string) (location.Location, error) {
	log.Info("Пользователь зашел с данными", "id", id, "location", newLocation)

	var exists bool
	loc := location.Location{}

	err := t.bd.QueryRow(context.Background(), checkUserInfo, id).Scan(&exists)
	if err != nil {
		log.Error("QueryRow failed cheack user info witdh id = ", id, " ", err)
		return location.Location{}, err
	}

	if !exists {
		err = t.bd.QueryRow(context.Background(), addLocation, id, newLocation).Scan(&loc.Id, &loc.Location)
		if err != nil {
			log.Error("Failed add row in user table with id", "id", id, "error", err)
			return location.Location{}, err
		}
		return loc, nil
	}

	err = t.bd.QueryRow(context.Background(), updateLocation, newLocation, id).Scan(&loc.Id, &loc.Location)
	if err != nil {
		log.Error("Failed update row in user table with id", "id", id, "error", err)
		return location.Location{}, err
	}
	return loc, nil
}

func (t *Transport) GetLocationInfo(id int64) (location.Location, error) {
	log.Info("Пользователь зашел с данными", "id", id)

	var exists bool
	loc := location.Location{}

	err := t.bd.QueryRow(context.Background(), checkUserInfo, id).Scan(&exists)
	if err != nil {
		log.Error("QueryRow failed cheack user info with id = ", id, " ", err)
		return location.Location{}, err
	}

	if !exists {
		return location.Location{}, errors.New("У пользователя не установлен город")
	}

	err = t.bd.QueryRow(context.Background(), getUserInfo, id).Scan(&loc.Id, &loc.Location)
	if err != nil {
		log.Error("Failed get row in user table with id", "id", id, "error", err)
		return location.Location{}, err
	}
	return loc, nil
}
