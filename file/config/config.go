package config

import "github.com/caarlos0/env/v9"

var Config struct {
	Mode  string `env:"MODE" envDefault:"dev"`
	Debug bool   `env:"DEBUG" envDefault:"false"`
}

func Init() {
	err := env.Parse(&Config)
	if err != nil {
		panic(err)
	}
}
