package config

import (
	"log"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	APIToken       string        `env:"API_TOKEN,required"`
	OwnerSlug      string        `env:"OWNER_SLUG,required"`
	MongoURI       string        `env:"MONGO_URI,required"`
	DefaultTimeout time.Duration `env:"DEFAULT_TIMEOUT,required"`
	DBName         string        `env:"DB_NAME,required"`
}

func Parse() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
		return Config{}
	}

	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal("unable to parse environment variables: ", err)
		return Config{}
	}

	return cfg
}
