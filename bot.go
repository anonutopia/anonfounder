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
	} else if strings.HasPrefix(tu.Message.Text, "/freeinfo") {
		freeinfoCommand(tu)
	} else if strings.HasPrefix(tu.Message.Text, "/") {
		unknownCommand(tu)
	} else if tu.UpdateID != 0 {
		log.Println(tu.Message.ReplyToMessage.MessageID)
		if tu.Message.ReplyToMessage.MessageID == 0 {
			if tu.Message.NewChatMember.ID != 0 {
				messageTelegram(fmt.Sprintf("Welcome %s! 🚀", tu.Message.NewChatMember.FirstName), int64(tu.Message.Chat.ID))
			}
		} else {
			avr, err := wnc.AddressValidate(tu.Message.Text)
			if err != nil {
				logTelegram(err.Error())
				messageTelegram("Something went wrong, please try again.", int64(tu.Message.Chat.ID))
			} else {
				if !avr.Valid {
					messageTelegram("Your wallet address is not valid. Please check if it's correct and try again.", int64(tu.Message.Chat.ID))
				} else {
					tu.Message.Text = fmt.Sprintf("/drop %s", tu.Message.Text)
					dropCommand(tu)
				}
			}
		}
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
	pc.Caption = "QR Code"
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
	kv := &KeyValue{Key: "airdropSent"}
	db.First(kv, kv)
	if kv.ValueInt >= conf.Airdrop {
		messageTelegram("We are sorry to inform you that token airdrop has already finished.", int64(tu.Message.Chat.ID))
		return
	}

	msgArr := strings.Fields(tu.Message.Text)
	if len(msgArr) == 1 {
		msg := tgbotapi.NewMessage(int64(tu.Message.Chat.ID), "Please enter your Waves address")
		msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, Selective: true}
		msg.ReplyToMessageID = tu.Message.MessageID
		bot.Send(msg)
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

							kv.ValueInt = kv.ValueInt + 1
							db.Save(kv)

							messageTelegram("I've sent you your 1 free token. Welcome! 🚀", int64(tu.Message.Chat.ID))
						}
					}
				}
			}
		}
	}
}

func freeinfoCommand(tu TelegramUpdate) {
	kv := &KeyValue{Key: "airdropSent"}
	db.First(kv, kv)
	msg := fmt.Sprintf("Tokens dropped so far: <strong>%d ATST</strong>\nTokens left to drop: <strong>%d ATST</strong>", kv.ValueInt, (conf.Airdrop - kv.ValueInt))
	log.Println(msg)
	m := tgbotapi.NewMessage(int64(tu.Message.Chat.ID), msg)
	m.ParseMode = "HTML"
	err, e := bot.Send(m)
	log.Println(err)
	log.Println(e)
}

func unknownCommand(tu TelegramUpdate) {
	messageTelegram("This command doesn't exist.", int64(tu.Message.Chat.ID))
}
