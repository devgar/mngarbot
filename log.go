package main

import (
	"fmt"
	"log"

	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

func logMsgData(m tb.Message) {
	mFromID := ""
	mFromName := ""
	mType := "Chat"
	if m.From != nil {
		mFromID = fmt.Sprintf("%d", m.From.ID)
		mFromName = m.From.UserName
	}
	if m.Chat.IsChannel() {
		mType = "Channel"
	}
	if m.Chat.IsGroup() {
		mType = "Group"
	}
	if m.Chat.IsSuperGroup() {
		mType = "SuperGroup"
	}
	log.Printf(
		"(%s)\n#%s [%s]\n#%d [%s]\n  %s",
		mType,
		mFromID,
		mFromName,
		m.Chat.ID,
		m.Chat.Title,
		m.Text,
	)
}
