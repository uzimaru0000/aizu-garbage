package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MySQL MySQL `envconfig:"MYSQL"`
}

type MySQL struct {
	User     string
	Password string
	Name     string `envconfig:"DATABASE"`
	Host     string
}

var config Config

func Init(envPath string) {
	mode := os.Getenv("MODE")
	if mode == "DEV" {
		err := godotenv.Load(envPath)
		if err != nil {
			panic(err)
		}
	}

	err := envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &config
}
