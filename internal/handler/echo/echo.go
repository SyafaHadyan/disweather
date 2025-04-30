package echo

import (
	"strings"

	"github.com/SyafaHadyan/disweather/internal/handler/interactionrespond"
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

	interactionrespond.InteractionRespond(s, i, builder.String(), "echo")
}
