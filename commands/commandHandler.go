package commands

import (
	"log"
	"malu/commands/cmds"
	"malu/config"

	"github.com/bwmarrin/discordgo"
)

// Commands Guarda os comandos
var (
	Commands = make(map[string]Command)
)

// RegisterCommands Registra os comandos.
func RegisterCommands() {
	Commands["ping"] = Command{cmds.PingCommand, false}
	Commands["calc"] = Command{cmds.CalcCommand, false}
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

	return
}
