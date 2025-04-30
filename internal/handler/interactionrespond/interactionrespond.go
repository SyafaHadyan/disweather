package interactionrespond

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func InteractionRespond(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	message string,
	origin string,
) {
	log.Printf("[COMMAND] %s: %#v\n", origin, message)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
	if err != nil {
		log.Panicf("[COMMAND] %s: Could not respond to interaction: %s", origin, err)
	}
}
