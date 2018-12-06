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
		SetDescription("512x512 PNG").
		SetImage(user.AvatarURL("512"))

	s.ChannelMessageSendEmbed(c.ID, embed.MessageEmbed)
	return nil
}
