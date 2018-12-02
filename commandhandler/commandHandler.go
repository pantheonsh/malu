package commandhandler

import (
	"log"
	"malu/config"
	"malu/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var aliases = make(map[string]Command)

// RegisterAliases Adiciona os nomes de aliases como referências aos comandos.
func RegisterAliases() {
	for _, cmd := range Commands {
		for _, aliasName := range cmd.Aliases {
			aliases[aliasName] = cmd
		}
	}
}

// ExecCommand Executa um comando a partir do seu nome.
func ExecCommand(commandName string, args []string, channel *discordgo.Channel, message *discordgo.MessageCreate, session *discordgo.Session) {
	// Retorna o comando se ele existe
	if cmd, found := FindCommand(commandName); found {
		if cmd.OwnerOnly && message.Author.ID != config.Data.Owner {
			session.ChannelMessageSend(channel.ID, "Este comando só é acessível aos desenvolvedores.")
			return
		}

		err := cmd.Exec(args, session, message, channel)

		if err != nil {
			log.Println("Erro ao executar o comando", commandName, args, "no canal", channel.Name)
			log.Println(err.Error())

			session.ChannelMessageSend(channel.ID, "Ocorreu um erro ao executar o comando!")
		}
	}

	if commandName == "ajuda" {
		HelpSpecialCommand(args, session, message, channel)
	}
}

// FindCommand Encontra um comando pelo seu nome ou por um alias
func FindCommand(cmdname string) (Command, bool) {
	var c = Command{}

	if cmd, comandoExiste := Commands[cmdname]; comandoExiste {
		return cmd, true
	}

	for _, cmd := range Commands {
		for _, alias := range cmd.Aliases {
			if alias == cmdname {
				return cmd, true
			}
		}
	}

	return c, false
}

// HelpSpecialCommand Comando especial de ajuda
func HelpSpecialCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) {
	var embed = utils.NewEmbed().
		SetTitle("Comandos").
		SetColor(0x7289DA)

	for _, val := range Commands {
		embed.AddField("*"+val.Example+"*", val.Description+"\nOutros nomes: "+strings.Join(val.Aliases, ", "))
	}

	s.ChannelMessageSendEmbed(c.ID, embed.MessageEmbed)
}
