package main

import (
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	tAnonBalkan = -1001161265502
	tAnon       = -1001361489843
	tAnonTaxi   = -1001422544298
	tAnonTaxiP  = -1001271198034
	tAnonOps    = -297434742
)

func initBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(conf.TelegramAPIKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = conf.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(tAnonOps, "AnonFounder successfully started. 🚀")
	bot.Send(msg)

	return bot
}

func logTelegram(message string) {
	msg := tgbotapi.NewMessage(tAnonOps, message)
	bot.Send(msg)
}
