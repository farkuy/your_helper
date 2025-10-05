package messages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Message struct {
	UserId int64
	Text   string
}

type MessageSender interface {
	NewMessage(userId int64, text string) tgbotapi.MessageConfig
}

type Model struct {
	tgClient MessageSender
}

func Init(m *MessageSender) *Model {
	return &Model{tgClient: *m}
}

func (m *Model) PostAnswer(msg Message) {
	m.tgClient.NewMessage(msg.UserId, "Без понятия что ты от меня хочешь")
}
