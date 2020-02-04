package main

import (
	"log"

	macaron "gopkg.in/macaron.v1"
)

func pageView(ctx *macaron.Context) string {
	log.Println("fdsafdsa")
	return "OK"
}
