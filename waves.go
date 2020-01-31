package main

import (
	"log"

	"github.com/anonutopia/gowaves"
)

func initWaves() *gowaves.WavesNodeClient {
	log.Println(conf.WavesNodeAPIKey)
	wnc := &gowaves.WavesNodeClient{
		Host:   "anode1.anonutopia.com",
		Port:   6869,
		ApiKey: conf.WavesNodeAPIKey,
	}

	return wnc
}
