package main

import (
	"log"
	"your_helper/internal/bot"
	"your_helper/internal/config"
	log_wrapper "your_helper/internal/log"
	"your_helper/internal/model/messages"
)

// TODO: сделать запрос по погоде надень
func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	log_wrapper.Init(cfg.Environment)

	msg := messages.Init()
	tgBot := bot.Init(cfg.TgBotToken, msg)

	tgBot.Listener()
}
