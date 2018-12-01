package handlers

import (
	"malu/commands"
	"malu/config"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignoraremos mensagens de outros bots.
	if m.Author.Bot {
		return
	}

	channel, err := s.Channel(m.ChannelID)

	// Se não conseguirmos pegar o canal, ignoraremos antes que seja tarde.
	if err != nil {
		return
	}

	// Por enquanto, só queremos aceitar mensagens de canais de texto de servidores.
	if channel.Type != discordgo.ChannelTypeGuildText {
		return
	}

	var content = m.Content

	// Se a mensagem não começar com um prefixo, ignoraremos também.
	if strings.HasPrefix(content, config.Data.Prefix) != true {
		return
	}

	var split = strings.Split(content, " ")
	var cmd, args = split[0], split[1:]

	cmd = strings.TrimPrefix(cmd, config.Data.Prefix)
	commands.ExecCommand(cmd, args, channel, m, s)
}
