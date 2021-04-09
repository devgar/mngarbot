package main

import (
	"flag"
	"log"
	"os"

	tb "github.com/go-telegram-bot-api/telegram-bot-api"

	"edgarai.com/mngarbot/config"
)

var (
	chat    int64  = 0
	message string = ""
)

func init() {
	flag.Int64Var(&chat, "c", 0, "Chat to send a message")
	flag.StringVar(&message, "m", "", "Message to send to admin")
	flag.Parse()
}

func main() {
	c := config.Get()
	b, err := tb.NewBotAPI(c.Token)

	if err != nil {
		log.Panic(err)
	}

	if message != "" {
		if chat == 0 {
			chat = c.ID
		}
		// msg := tb.NewMessage(chat, message)
		// msg.ReplyMarkup = tb.NewRemoveKeyboard(true)
		msg := NewInlineNumericKeyboardMsg(chat, message)
		b.Send(msg)
		os.Exit(0)
	}

	if c.ID != 0 {
		msg := tb.NewMessage(c.ID, "MNGRBOT Started")
		b.Send(msg)
	}

	u := tb.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			message := update.Message
			if c.ID == 0 {
				log.Printf("[No ADMIN set] %d %s", message.From.ID, message.From.UserName)
				continue
			}

			if message.Chat.IsPrivate() && c.ID != int64(message.From.ID) {
				msg := tb.NewMessage(message.Chat.ID, "I'm not allowed to talk to you")
				b.Send(msg)
			}

			if message.Chat.IsPrivate() && c.ID == int64(message.From.ID) {
				if message.Text == "/sample" {
					msg := NewInlineNumericKeyboardMsg(message.Chat.ID, "uWu")
					b.Send(msg)
				}
			}

			logMsgData(*message)
		} else if update.ChannelPost != nil {
			logMsgData(*update.ChannelPost)
		} else if update.EditedMessage != nil {
			log.Printf("Edited message:\n  %s", update.EditedMessage.Text)
		} else if update.EditedChannelPost != nil {
			log.Printf("Edited message:\n  %s", update.EditedChannelPost.Text)
		} else if update.CallbackQuery != nil {
			callback := tb.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			// if _, err := b.Request(callback); err != nil {
			// 	panic(err)
			// }
			msg := tb.NewMessage(update.CallbackQuery.Message.Chat.ID, callback.Text)
			b.Send(msg)
		} else {
			log.Printf("Unhandled update type")
		}
	}
}
