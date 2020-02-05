package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/anonutopia/gowaves"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const satInBtc = uint64(100000000)

func executeBotCommand(tu TelegramUpdate) {
	if strings.HasPrefix(tu.Message.Text, "/price") {
		priceCommand(tu)
	} else if strings.HasPrefix(tu.Message.Text, "/start") {
		startCommand(tu)
	} else if strings.HasPrefix(tu.Message.Text, "/address") {
		addressCommand(tu)
	} else if strings.HasPrefix(tu.Message.Text, "/balance") {
		balanceCommand(tu)
	} else if strings.HasPrefix(tu.Message.Text, "/drop") {
		dropCommand(tu)
	} else {
		unknownCommand(tu)
	}
}

func priceCommand(tu TelegramUpdate) {
	kv := &KeyValue{Key: "currentPrice"}
	db.First(kv, kv)
	price := float64(kv.ValueInt) / float64(satInBtc)
	messageTelegram(fmt.Sprintf("Current token price is: %.8f €", price), int64(tu.Message.Chat.ID))
}

func startCommand(tu TelegramUpdate) {
	messageTelegram("Hello and welcome to Anonutopia!", int64(tu.Message.Chat.ID))
}

func addressCommand(tu TelegramUpdate) {
	messageTelegram("My main Waves address is:", int64(tu.Message.Chat.ID))
	messageTelegram(conf.NodeAddress, int64(tu.Message.Chat.ID))
	pc := tgbotapi.NewPhotoUpload(int64(tu.Message.Chat.ID), "qrcode.png")
	bot.Send(pc)
}

func balanceCommand(tu TelegramUpdate) {
	b, err := wnc.AddressesBalance(conf.NodeAddress)
	if err != nil {
		log.Printf("balanceCommand error: %s", err)
	}
	messageTelegram(fmt.Sprintf("My current Waves balance is: %.8f WAVES", float64(b.Balance)/float64(satInBtc)), int64(tu.Message.Chat.ID))
}

func dropCommand(tu TelegramUpdate) {
	msgArr := strings.Fields(tu.Message.Text)
	if len(msgArr) == 1 {
		messageTelegram("Wallet address is required. Please try again providing address this time (/drop address).", int64(tu.Message.Chat.ID))
	} else {
		avr, err := wnc.AddressValidate(msgArr[1])
		if err != nil {
			logTelegram(err.Error())
			messageTelegram("Something went wrong, please try again.", int64(tu.Message.Chat.ID))
		} else {
			if !avr.Valid {
				messageTelegram("Your wallet address is not valid. Please check if it's correct and try again.", int64(tu.Message.Chat.ID))
			} else {
				user := &User{TelegramID: tu.Message.From.ID}
				db.First(user, user)

				if user.ID != 0 {
					if user.Address == msgArr[1] {
						messageTelegram("Your free token is already activated. You'll have to be a better hacker than that! 😆", int64(tu.Message.Chat.ID))
					} else {
						messageTelegram("This is already becoming some serious hacking? Obviously not serious enough. 😎", int64(tu.Message.Chat.ID))
					}
				} else {
					if msgArr[1] == conf.NodeAddress {
						messageTelegram("You need to use your wallet address, the one you get with Waves wallet.", int64(tu.Message.Chat.ID))
					} else {
						atr := &gowaves.AssetsTransferRequest{
							Amount:    100000000,
							AssetID:   conf.TokenID,
							Fee:       100000,
							Recipient: msgArr[1],
							Sender:    conf.NodeAddress,
						}

						_, err := wnc.AssetsTransfer(atr)
						if err != nil {
							messageTelegram("Something went wrong, please try again.", int64(tu.Message.Chat.ID))
							logTelegram(err.Error())
						} else {
							user.TelegramID = tu.Message.From.ID
							user.TelegramUsername = tu.Message.From.Username
							user.Address = msgArr[1]
							db.Save(user)

							messageTelegram("I've sent you your 1 free token. Welcome! 🚀", int64(tu.Message.Chat.ID))
						}
					}
				}
			}
		}
	}
}

func unknownCommand(tu TelegramUpdate) {
	messageTelegram("This command doesn't exist.", int64(tu.Message.Chat.ID))
}
