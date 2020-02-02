package main

import (
	"github.com/anonutopia/gowaves"
	"github.com/jinzhu/gorm"
)

var conf *Config

var wnc *gowaves.WavesNodeClient

var db *gorm.DB

func main() {
	conf = initConfig()

	db = initDb()

	wnc = initWaves()

	initMonitor()
}
