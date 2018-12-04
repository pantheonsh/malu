package commands

import (
	"malu/utils"

	"github.com/bwmarrin/discordgo"
)

func AvatarCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error {
	var user = m.Author

	if len(m.Mentions) > 0 {
		user = m.Mentions[0]
	}

	var embed = utils.NewEmbed().
		SetColor(0x4286f4).
		SetTitle("Avatar de " + user.Username + "#" + user.Discriminator).
		SetImage(user.AvatarURL(""))

	s.ChannelMessageSendEmbed(c.ID, embed.MessageEmbed)
	return nil
}
