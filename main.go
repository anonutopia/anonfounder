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

	ab, _ := wnc.AddressesBalance("3PLJQASFXtiohqbirYwSswjjbYGLfaGDEQy")

	db.FirstOrCreate(&KeyValue{Key: "test", ValueInt: 25})

	log.Println(ab.Balance)
}
