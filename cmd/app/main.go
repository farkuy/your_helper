package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"your_helper/database"
	"your_helper/internal/bot"
	"your_helper/internal/config"
	log_wrapper "your_helper/internal/log"
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
	wetaherModel := weather.Init(&weatherTr)

	msg := messages.Init(*wetaherModel)
	tgBot := bot.Init(cfg.TgBotToken, msg)

	tgBot.Listener()
}
