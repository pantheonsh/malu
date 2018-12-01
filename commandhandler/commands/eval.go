package commands

import (
	"strings"

	"malu/utils"

	"github.com/bwmarrin/discordgo"
	"github.com/robertkrimen/otto"
)

var JSvm = otto.New()

func EvalCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error {
	code := strings.Join(args, " ")

	JSvm.Set("s", s)
	JSvm.Set("m", m)
	JSvm.Set("c", c)

	value, _ := JSvm.Run(code)
	str, _ := value.ToString()

	var embed = utils.NewEmbed().
		SetColor(0x4286f4).
		SetTitle("Interpretador JavaScript").
		AddField("Resultado", str)

	s.ChannelMessageSendEmbed(c.ID, embed.MessageEmbed)
	return nil
}
