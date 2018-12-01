package commands

import "github.com/bwmarrin/discordgo"

// Command A base para os comandos
type Command struct {
	Exec      func(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error
	OwnerOnly bool
}
