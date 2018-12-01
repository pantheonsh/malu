package cmds

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

import "gopkg.in/resty.v1"

func CalcCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error {
	expression := strings.Join(args, " ")
	reqbody := map[string]interface{}{
		"expr":      expression,
		"precision": 42,
	}

	var result map[string]interface{}
	var reserror map[string]interface{}

	_, err := resty.R().
		SetBody(reqbody).
		SetResult(&result).
		SetError(&reserror).
		Post("http://api.mathjs.org/v4/")

	if err != nil {
		return err
	}

	if reserror["error"] != nil {
		s.ChannelMessageSend(c.ID, reserror["error"].(string))
		return nil
	}

	s.ChannelMessageSend(c.ID, result["result"].(string))

	return nil
}
