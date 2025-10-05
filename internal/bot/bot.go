package bot

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	Api tgbotapi.BotAPI
}

func Init(token string) *TgBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		slog.Error("ошибка при подключении бота: %v", err)
	}
	slog.Info("Бот успешно получил соединение")

	return &TgBot{Api: *bot}
}

func (bot *TgBot) Listener() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.Api.GetUpdatesChan(u)

	slog.Info("Бот начал цикл прослушки")
	for update := range updates {
		if update.Message != nil {
			var msg tgbotapi.MessageConfig

			messages := update.Message.Text
			userId := update.Message.Chat.ID

			switch messages {
			case "/start":
				msg = tgbotapi.NewMessage(userId, "Жми екарный бабай")
			default:
				msg = tgbotapi.NewMessage(userId, "Я хз что это")
			}

			bot.Api.Send(msg)
		}
	}
}
