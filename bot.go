package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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
	} else {
		unknownCommand(tu)
	}
}

func priceCommand(tu TelegramUpdate) {
	kv := &KeyValue{Key: "currentPrice"}
	db.First(kv, kv)
	messageTelegram(fmt.Sprintf("Current token price is: %s €", strconv.Itoa(int(kv.ValueInt))), int64(tu.Message.Chat.ID))
}

func startCommand(tu TelegramUpdate) {
	messageTelegram("Hello and welcome to Anonutopia!", int64(tu.Message.Chat.ID))
}

func addressCommand(tu TelegramUpdate) {
	messageTelegram("My main Waves address is:", int64(tu.Message.Chat.ID))
	messageTelegram(conf.NodeAddress, int64(tu.Message.Chat.ID))
}

func balanceCommand(tu TelegramUpdate) {
	b, err := wnc.AddressesBalance(conf.NodeAddress)
	if err != nil {
		log.Printf("balanceCommand error: %s", err)
	}
	messageTelegram(fmt.Sprintf("My current Waves balance is: %.3f WAVES", float64(b.Balance)/float64(satInBtc)), int64(tu.Message.Chat.ID))
}

func unknownCommand(tu TelegramUpdate) {
	messageTelegram("This command doesn't exist.", int64(tu.Message.Chat.ID))
}
