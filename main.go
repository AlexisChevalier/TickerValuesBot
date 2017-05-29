package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
	"./api"
	"os"
)

type Configuration struct {
	token string
}

func main() {

	token := os.Args[1]
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		text := update.Message.Text

		if len(text) == 6 {
			from := text[:3]
			to := text[3:]

			result := api.GetTickerValue(from, to)

			if result == nil || !result.Success {
				response := fmt.Sprintf("Conversion from %s to %s failed", from, to)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				bot.Send(msg)
			} else {
				response := fmt.Sprintf("%d %s = %s %s", 1, result.Ticker.Base, result.Ticker.Price, result.Ticker.Target)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				bot.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid tickers")
			bot.Send(msg)
		}
	}
}