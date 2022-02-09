package config

import (
	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
)

// UrlChangerConfig config for chat BE
type UrlChangerConfig struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	DBPort   int    `env:"DB_PORT"`
	DBUser   string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
}

func GetConfig(configType interface{}) (err error) {
	godotenv.Load(".env")
	return gonfig.GetConf("", configType)
}
