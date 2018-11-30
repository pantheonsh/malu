package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config Configuração do bot
type Config struct {
	Token  string
	Prefix string
	Owner  string
}

// Data A configuração carregada
var (
	Data *Config
)

func Load() {
	// First, we load the config file to memory
	b, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = json.Unmarshal(b, &Data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
