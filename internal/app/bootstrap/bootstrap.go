package bootstrap

import (
	"log"

	"github.com/SyafaHadyan/disweather/internal/app/bot"
	"github.com/SyafaHadyan/disweather/internal/infra/env"
	"github.com/bwmarrin/discordgo"
)

func Start() {
	config, err := env.New()
	if err != nil {
		log.Fatal(err)
	}

	session, _ := discordgo.New("Bot " + config.Token)

	bot.NewBot(
		config,
		session,
	)
}
