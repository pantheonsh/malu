package cmds

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CalcCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error {
	expression := strings.Join(args, " ")
	reqbody := map[string]interface{}{
		"expr":      expression,
		"precision": 42,
	}

	reqbodybytes, err := json.Marshal(reqbody)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://api.mathjs.org/v4/", "application/json", bytes.NewBuffer(reqbodybytes))
	if err != nil {
		return err
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["error"] != nil {
		s.ChannelMessageSend(c.ID, result["error"].(string))
		return nil
	}

	s.ChannelMessageSend(c.ID, result["result"].(string))

	return nil
}
