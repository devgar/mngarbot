package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	tb "gopkg.in/tucnak/telebot.v2"

	"edgarai.com/mngarbot/config"
)

var TOKEN string
var message string = ""

func init() {
	TOKEN = os.Getenv("TOKEN")
	flag.StringVar(&message, "m", "", "Message to send to admin")
	flag.Parse()
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

	if message != "" {
		b.Send(&tb.User{ID: c.ID}, message)
		os.Exit(0)
	}

	if c.ID != 0 {
		b.Send(&tb.User{ID: c.ID}, "Starting MNGRBOT service")
	}

	b.Handle("/hello", func(m *tb.Message) {
		log.Printf("[HELLO] USER ID: %d", m.Sender.ID)
		b.Send(m.Sender, "Hello World!")
	})

	b.Start()
}
