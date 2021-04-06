package main

import (
	"flag"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"

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
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Printf("ADMIN ID: %d\n", c.ID)
	// }
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
