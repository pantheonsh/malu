package commandhandler

import (
	"log"
	"malu/commandhandler/commands"
	"malu/config"

	"github.com/bwmarrin/discordgo"
)

// Commands Guarda os comandos
var (
	Commands = make(map[string]Command)
)

// RegisterCommands Registra os comandos.
func RegisterCommands() {
	Commands["calc"] = Command{commands.CalcCommand, false}
	Commands["eval"] = Command{commands.EvalCommand, true}
	Commands["google"] = Command{commands.GoogleCommand, false}
}

// ExecCommand Executa um comando a partir do seu nome.
func ExecCommand(commandName string, args []string, channel *discordgo.Channel, message *discordgo.MessageCreate, session *discordgo.Session) {
	// Retorna o comando se ele existe no Map
	if cmd, comandoExiste := Commands[commandName]; comandoExiste {
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

// HelpSpecialCommand Comando especial de ajuda
func HelpSpecialCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) {
	var str = ""

	for key := range Commands {
		str += key + "\n"
	}

	s.ChannelMessageSend(c.ID, str)
}
