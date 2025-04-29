package command

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "echo",
		Description: "Say something through a bot",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "message",
				Description: "Contents of the message",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "author",
				Description: "Whether to prepend message's author",
				Type:        discordgo.ApplicationCommandOptionBoolean,
			},
		},
	},
	{
		Name:        "weather",
		Description: "Get weather",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "query",
				Description: "Get weather from any query",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
}
