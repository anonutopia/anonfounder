package main

import (
	"fmt"
	"strconv"
	"strings"
)

func executeBotCommand(tu TelegramUpdate) {
	if strings.HasPrefix(tu.Message.Text, "/price") {
		priceCommand(tu)
	} else if strings.HasPrefix(tu.Message.Text, "/start") {
		startCommand(tu)
	} else if strings.HasPrefix(tu.Message.Text, "/address") {
		addressCommand(tu)
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

func unknownCommand(tu TelegramUpdate) {
	messageTelegram("This command doesn't exist.", int64(tu.Message.Chat.ID))
}
