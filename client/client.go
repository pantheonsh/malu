package client

import (
	"log"
	"malu/client/handlers"
	"malu/commands"
	"malu/config"

	"github.com/bwmarrin/discordgo"
)

// Session Sessão do Discord atual
var (
	Session *discordgo.Session
)

// Start Inicia a conexão com o servidor Discord.
func Start() {
	config.Load()
	s, err := discordgo.New("Bot " + config.Data.Token)

	if err != nil {
		log.Fatalln(err.Error())
	}

	Session = s

	err = Session.Open()

	if err != nil {
		log.Fatalln(err.Error())
	}

	commands.RegisterCommands()

	Session.AddHandler(handlers.MessageCreate)

	log.Println("Iniciado!")
}
