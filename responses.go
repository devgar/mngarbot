package main

import (
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

var numericKeyboard = tb.NewReplyKeyboard(
	tb.NewKeyboardButtonRow(
		tb.NewKeyboardButton("1"),
		tb.NewKeyboardButton("2"),
		tb.NewKeyboardButton("3"),
	),
	tb.NewKeyboardButtonRow(
		tb.NewKeyboardButton("4"),
		tb.NewKeyboardButton("5"),
		tb.NewKeyboardButton("6"),
	),
	tb.NewKeyboardButtonRow(
		tb.NewKeyboardButton("7"),
		tb.NewKeyboardButton("8"),
		tb.NewKeyboardButton("9"),
	),
	tb.NewKeyboardButtonRow(
		tb.NewKeyboardButton("0"),
		tb.NewKeyboardButton("."),
	),
)

func NewNumericKeyboardMsg(id int64, text string) tb.MessageConfig {
	msg := tb.NewMessage(id, text)
	msg.ReplyMarkup = numericKeyboard
	return msg
}

func NewEchoMsg(m tb.Message) tb.MessageConfig {
	msg := tb.NewMessage(m.Chat.ID, m.Text)
	msg.ReplyToMessageID = m.MessageID
	return msg
}
