package config

type Config struct {
	Token  string
	Prefix string
	Owner  string
}

var (
	Data *Config
)

func Load() {

}
