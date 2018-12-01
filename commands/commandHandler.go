package commands

import (
	"log"
	"malu/commands/cmds"

	"github.com/bwmarrin/discordgo"
)

// Commands Guarda os comandos
var (
	Commands = make(map[string]Command)
)

// RegisterCommands Registra os comandos.
func RegisterCommands() {
	Commands["ping"] = Command{cmds.PingCommand, false}
}

// ExecCommand Executa um comando a partir do seu nome.
func ExecCommand(commandName string, args []string, channel *discordgo.Channel, message *discordgo.MessageCreate, session *discordgo.Session) {
	// Retorna o comando se ele existe no Map
	if cmd, comandoExiste := Commands[commandName]; comandoExiste {
		err := cmd.Exec(session, message, channel)

		if err != nil {
			log.Println("Erro ao executar o comando", commandName, args, "no canal", channel.Name)
			log.Println(err.Error())

			session.ChannelMessageSend(channel.ID, "Ocorreu um erro ao executar o comando!")
		}
	}

	return
}
