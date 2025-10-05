package main

import (
	"log"
	"your_helper/internal/bot"
	"your_helper/internal/config"
	log_wrapper "your_helper/internal/log"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	log_wrapper.Init(cfg.Environment)

	tgBot := bot.Init(cfg.TgBotToken)
	tgBot.Listener()
}
