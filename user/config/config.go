package config

import (
	"fmt"
	"github.com/caarlos0/env/v9"
)

var Config struct {
	Mode   string `env:"MODE" envDefault:"dev"`
	Debug  bool   `env:"DEBUG" envDefault:"false"`
	DbType string `env:"DB_TYPE" envDefault:"mysql"`
	DbURL  string `env:"DB_URL"`
}

func Init() {
	err := env.Parse(&Config)
	if err != nil {
		panic(err)
	}

	fmt.Println(Config)
}
