package main

import (
	"github.com/anonutopia/gowaves"
	"github.com/jinzhu/gorm"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var conf *Config

var wnc *gowaves.WavesNodeClient

var db *gorm.DB

var bot *tgbotapi.BotAPI

func main() {
	conf = initConfig()

	db = initDb()

	wnc = initWaves()

	bot = initBot()

	initMonitor()
}
