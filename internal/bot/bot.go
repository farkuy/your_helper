package bot

import (
	"log"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type RequestMsg interface {
	CreateAnswer(userId int64, text string) (string, error)
}

type TgBot struct {
	Api        tgbotapi.BotAPI
	MsgHandler RequestMsg
}

func Init(token string, msgHandler RequestMsg) *TgBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		slog.Error("ошибка при подключении бота: %v", err)
		log.Fatal(err)
	}
	slog.Info("Бот успешно получил соединение")

	return &TgBot{Api: *bot, MsgHandler: msgHandler}
}

func (bot *TgBot) SendMessage(userId int64, text string) {
	msg, err := bot.MsgHandler.CreateAnswer(userId, text)
	if err != nil {
		slog.Error("произошла ошибка: %v", err)
		msg = "Произошла ошибка обработки ответа"
	}

	msgToSend := tgbotapi.NewMessage(userId, msg)
	if _, err := bot.Api.Send(msgToSend); err != nil {
		slog.Error("Ошибка при отправке сообщения: %v", err)
	}
}

func (bot *TgBot) Listener() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.Api.GetUpdatesChan(u)

	slog.Info("Бот начал цикл прослушки")
	for update := range updates {
		if update.Message != nil {

			messages := update.Message.Text
			userId := update.Message.Chat.ID

			bot.SendMessage(userId, messages)
		}
	}
}
