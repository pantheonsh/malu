package commands

import (
	"malu/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func GoogleCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error {
	searchQuery := strings.Join(args, " ")
	links, err := utils.GoogleSearch(searchQuery)

	if err != nil {
		return err
	}

	if len(links) < 1 {
		s.ChannelMessageSend(c.ID, "Nenhum resultado encontrado.")
		return nil
	}

	s.ChannelMessageSend(c.ID, links[0])

	return nil
}
