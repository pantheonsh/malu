package cmds

import (
	"github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error {
	s.ChannelMessageSend(c.ID, "Pong!")
	return nil
}
