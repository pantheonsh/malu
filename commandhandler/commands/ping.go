package commands

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func PingCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error {
	msg, err := s.ChannelMessageSend(c.ID, "Aguarde.")

	if err != nil {
		return err
	}

	msgtimestamp, err := msg.Timestamp.Parse()

	if err != nil {
		return err
	}

	var t1 = msgtimestamp.Unix()
	var t2 = time.Now().Unix()

	s.ChannelMessageEdit(msg.ChannelID, msg.ID, "Ping: "+strconv.FormatInt((t2-t1)*2, 10)+"ms")

	return nil
}
