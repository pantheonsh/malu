package commandhandler

import "github.com/bwmarrin/discordgo"
import "malu/commandhandler/commands"

// Command A base para os comandos
type Command struct {
	Exec        func(args []string, s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel) error
	Name        string
	Aliases     []string
	Description string
	Example     string
	OwnerOnly   bool
}

// Commands Lista todos os comandos
var Commands = map[string]Command{
	"calc": Command{
		Exec:        commands.CalcCommand,
		Name:        "calc",
		Aliases:     []string{"math"},
		Description: "Calcula uma expressão matemática",
		Example:     "calc 1 + 1",
		OwnerOnly:   false,
	},
	"google": Command{
		Exec:        commands.GoogleCommand,
		Name:        "google",
		Aliases:     []string{"g"},
		Description: "Faz uma pesquisa no Google e retorna o primeiro resultado",
		Example:     "google Never gonna give you up",
		OwnerOnly:   false,
	},
	"eval": Command{
		Exec:        commands.EvalCommand,
		Name:        "eval",
		Aliases:     []string{"js"},
		Description: "Executa uma expressão JavaScript.",
		Example:     "eval console.log('1, 2, 3')",
		OwnerOnly:   true,
	},
}
