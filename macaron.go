package main

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/i18n"
	_ "github.com/go-macaron/session/redis"
	"github.com/mholt/certmagic"
	macaron "gopkg.in/macaron.v1"
)

func initMacaron() *macaron.Macaron {
	m := macaron.Classic()

	m.Use(cache.Cacher())

	m.Use(i18n.I18n(i18n.Options{
		Langs: []string{"hr", "sr", "en-US"},
		Names: []string{"Hrvatski", "Srpski", "English"},
	}))

	if conf.SSL {
		certmagic.Default.Agreed = true
		certmagic.Default.Email = "cryptopragmatic@protonmail.com"
		go certmagic.HTTPS([]string{"anonfounder.anonutopia.com"}, m)
	} else {
		go m.Run("0.0.0.0", 5000)
	}

	return m
}
