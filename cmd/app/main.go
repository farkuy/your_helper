package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"your_helper/database"
	"your_helper/internal/bot"
	"your_helper/internal/config"
	location_bd "your_helper/internal/database/location"
	log_wrapper "your_helper/internal/log"
	"your_helper/internal/models/location"
	"your_helper/internal/models/messages"
	"your_helper/internal/models/weather"
	weather_tranport "your_helper/internal/transports/rest/weather"
)

// TODO: сделать запрос по погоде надень
func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	log_wrapper.Init(cfg.Environment)
	db, err := database.Init(cfg.BdConfig)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close(context.Background())

	weatherTr := weather_tranport.Init(cfg.WeaterToken)
	wetaherModel := weather.Init(weatherTr)

	locationBd := location_bd.Init(db)
	locationModel := location.Init(&locationBd)

	msg := messages.Init(wetaherModel, locationModel)
	tgBot := bot.Init(cfg.TgBotToken, msg)

	tgBot.Listener()
}
