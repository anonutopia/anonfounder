package main

import (
	"fmt"
	"strconv"

	macaron "gopkg.in/macaron.v1"
)

func pageView(ctx *macaron.Context, tu TelegramUpdate) string {
	kv := &KeyValue{Key: "currentPrice"}
	db.First(kv, kv)
	messageTelegram(fmt.Sprintf("%s €", strconv.Itoa(int(kv.ValueInt))), int64(tu.Message.Chat.ID))

	return "OK"
}
