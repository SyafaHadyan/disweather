package bootstrap

import (
	"log"

	"github.com/SyafaHadyan/disweather/internal/app/bot"
	"github.com/SyafaHadyan/disweather/internal/infra/env"
	"github.com/bwmarrin/discordgo"
)

func Start() {
	log.Println("Started app")

	config, err := env.New()
	if err != nil {
		log.Fatal(err)
	}

	if config.Token == "" {
		log.Fatal("Token is not set")
	}

	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("Could not construct Discord client: %s", err)
	}

	log.Println("Successfully constructed client")

	bot.NewBot(
		config,
		session,
	)
}
