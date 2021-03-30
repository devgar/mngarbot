package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"edgarai.com/mngarbot/config"

	_ "github.com/joho/godotenv/autoload"
	tb "gopkg.in/tucnak/telebot.v2"
)

var TOKEN string

func init() {
	TOKEN = os.Getenv("TOKEN")
}

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("ADMIN ID: %d\n", c.ID)
	}
	b, err := tb.NewBot(tb.Settings{
		// URL: "https://telebot.edgarai.com",
		Token:  c.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		log.Printf("[HELLO] USER ID: %d", m.Sender.ID)
		b.Send(m.Sender, "Hello World!")
	})

	b.Start()
}
