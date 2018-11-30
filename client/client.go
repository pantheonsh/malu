package client

import (
	"log"
	"malu/config"

	"github.com/bwmarrin/discordgo"
)

var (
	Session *discordgo.Session
)

// Start Inicia a conexão com o servidor Discord.
func Start() {
	config.Load()
	s, err := discordgo.New("Bot" + config.Data.Token)

	if err != nil {
		log.Fatalln(err.Error())
	}

	Session = s

	err = Session.Open()

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Iniciado!")
}
