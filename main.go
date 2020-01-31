package main

import (
	"log"

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

	ab, _ := wnc.AddressesBalance(conf.NodeAddress)

	kv := &KeyValue{Key: "test", ValueInt: 65}

	db.FirstOrCreate(kv)

	kv.ValueInt = 65

	db.Save(kv)

	log.Println(ab.Balance)
}
