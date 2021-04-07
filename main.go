package main

import (
	"flag"
	"log"
	"os"

	tb "github.com/go-telegram-bot-api/telegram-bot-api"

	"edgarai.com/mngarbot/config"
)

var (
	message string = ""
)

func init() {
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
		msg := tb.NewMessage(c.ID, message)
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
		if update.Message == nil {
			log.Printf("Ignored empty message")
		}

		if c.ID == 0 {
			log.Printf("[Not admin from] %d %s", update.Message.From.ID, update.Message.From.UserName)
			continue
		}

		if c.ID != int64(update.Message.From.ID) {
			msg := tb.NewMessage(update.Message.Chat.ID, "I'm not allowed to talk to you")
			b.Send(msg)
		}

		// update

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tb.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		b.Send(msg)
	}
}
