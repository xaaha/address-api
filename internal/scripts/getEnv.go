package scripts

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ENV contains all the env var availale in .env file
type ENV struct {
	APIKey string
	DBPath string
	Port   string
}

// GetEnv gets the environment var set in .env
func GetEnv() ENV {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return ENV{
		APIKey: os.Getenv("API_KEY"),
		DBPath: os.Getenv("DB_FILE_PATH"),
		Port:   os.Getenv("PORT"),
	}
}
