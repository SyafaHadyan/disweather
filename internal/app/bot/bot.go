package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/SyafaHadyan/disweather/internal/domain/command"
	"github.com/SyafaHadyan/disweather/internal/handler/echo"
	"github.com/SyafaHadyan/disweather/internal/handler/weather"
	"github.com/SyafaHadyan/disweather/internal/infra/env"
	"github.com/bwmarrin/discordgo"
)

type optionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

type Bot struct {
	Config  *env.Env
	Session *discordgo.Session
}

func NewBot(config *env.Env, session *discordgo.Session) {
	b := Bot{
		Config:  config,
		Session: session,
	}

	b.Start()
}

func (b *Bot) parseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) optionMap {
	opMap := make(optionMap)
	for _, opt := range options {
		opMap[opt.Name] = opt
	}

	return opMap
}

func (b *Bot) Start() {
	if b.Config.ApplicationID == "" {
		log.Fatal("Application id is not set")
	}

	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()

		opMap := b.parseOptions(data.Options)

		switch data.Name {
		case "echo":
			echo.HandleEcho(s, i, opMap)
		case "weather":
			weather.HandleWeahter(s, i, opMap, b.Config.OpenWeatherAPI)
		default:
			return
		}
	})

	b.Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
	})

	_, err := b.Session.ApplicationCommandBulkOverwrite(
		b.Config.ApplicationID,
		b.Config.GuildID,
		command.Commands)
	if err != nil {
		log.Fatalf("Could not register commands: %s", err)
	}

	err = b.Session.Open()
	if err != nil {
		log.Fatalf("Could not open session: %s", err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = b.Session.Close()
	if err != nil {
		log.Printf("Could not close session gracefully: %s", err)
	}

	log.Println("Successfully closed session")
}
