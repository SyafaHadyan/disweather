package env

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Env struct {
	Token          string `env:"TOKEN"`
	ApplicationID  string `env:"APPLICATION_ID"`
	GuildID        string `env:"GUILD_ID"`
	OpenWeatherAPI string `env:"OPENWEATHER_API"`
}

func New() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	envParse := new(Env)
	err = env.Parse(envParse)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return envParse, nil
}
