package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Db *Db
}

type Db struct {
	Host     string
	User     string
	Password string
	Keyspace string
}

var once sync.Once
var configInstance *Config

func InitConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		configInstance = &Config{
			Db: &Db{
				Host:     os.Getenv("DB_HOST"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASS"),
				Keyspace: os.Getenv("DB_KEYSPACE"),
			},
		}
	})

	return configInstance
}
