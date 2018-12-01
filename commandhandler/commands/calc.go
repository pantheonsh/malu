package commands

import (
	"malu/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
	resty "gopkg.in/resty.v1"
)

func CalcCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error {
	expression := strings.Join(args, " ")
	reqbody := map[string]interface{}{
		"expr":      expression,
		"precision": 42,
	}

	var result map[string]interface{}
	var reserror map[string]interface{}
	var success = false
	var resstring string
	var embed = utils.NewEmbed()

	resty.R().
		SetBody(reqbody).
		SetResult(&result).
		SetError(&reserror).
		Post("http://api.mathjs.org/v4/")

	if reserror["error"] != nil {
		resstring = reserror["error"].(string)
	} else {
		success = true
		resstring = result["result"].(string)
	}

	embed.SetColor(0x4286f4).
		SetTitle("Calculadora").
		AddField("Expressão", expression).
		AddField("Resultado", resstring)

	if success != true {
		embed.SetColor(0xef3e50)
	}

	s.ChannelMessageSendEmbed(c.ID, embed.MessageEmbed)
	return nil
}
