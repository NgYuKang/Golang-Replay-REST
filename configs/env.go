package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file, exiting...")
	}
}

func EnvGinMode() string {
	mode := os.Getenv("GIN_MODE")

	if mode != "debug" && mode != "release" {
		return "release"
	}
	return mode
}

func EnvDBSource() string {
	dbSource := os.Getenv("DB_SOURCE")
	return dbSource

}
