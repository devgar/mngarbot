package main

import (
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

func echo(m tb.Message) tb.MessageConfig {
	msg := tb.NewMessage(m.Chat.ID, m.Text)
	msg.ReplyToMessageID = m.MessageID
	return msg
}
