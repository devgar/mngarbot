package main

import (
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

var inlineNumericKeyboard = tb.NewInlineKeyboardMarkup(
	tb.NewInlineKeyboardRow(
		tb.NewInlineKeyboardButtonURL("1.com", "http://ledro.es"),
		tb.NewInlineKeyboardButtonData("2", "o - 2"),
		tb.NewInlineKeyboardButtonData("3", "o - 3"),
	),
	tb.NewInlineKeyboardRow(
		tb.NewInlineKeyboardButtonData("4", "o - 4"),
		tb.NewInlineKeyboardButtonData("5", "5"),
		tb.NewInlineKeyboardButtonData("6", "6"),
	),
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

func NewInlineNumericKeyboardMsg(id int64, text string) tb.MessageConfig {
	msg := tb.NewMessage(id, text)
	msg.ReplyMarkup = inlineNumericKeyboard
	return msg
}

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
