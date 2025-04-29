package echo

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func interactionAuthor(i *discordgo.Interaction) *discordgo.User {
	if i.Member != nil {
		return i.Member.User
	}
	return i.User
}

func HandleEcho(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	opts map[string]*discordgo.ApplicationCommandInteractionDataOption,
) {
	builder := new(strings.Builder)
	if v, ok := opts["author"]; ok && v.BoolValue() {
		author := interactionAuthor(i.Interaction)
		builder.WriteString("**" + author.String() + "** says: ")
	}
	builder.WriteString(opts["message"].StringValue())

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: builder.String(),
		},
	})
	if err != nil {
		log.Panicf("could not respond to interaction: %s", err)
	}
}
