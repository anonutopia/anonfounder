package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "anonfounder.db")

	if err != nil {
		log.Printf("[initDb] error: %s", err)
	}

	db.DB()
	db.DB().Ping()
	db.LogMode(conf.Debug)
	db.AutoMigrate(&KeyValue{}, &Transaction{}, &User{})

	ks := &KeyValue{Key: "airdropSent", ValueInt: 0}
	db.FirstOrCreate(ks, ks)

	return db
}
