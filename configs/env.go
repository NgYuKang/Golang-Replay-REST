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

func EnvAWSRegion() string {
	ret := os.Getenv("AWS_S3_REGION")
	return ret

}

func EnvAWSBucket() string {
	ret := os.Getenv("AWS_S3_BUCKET")
	return ret

}

func EnvEncryptKey() string {
	ret := os.Getenv("ENCRYPT_SECRET_KEY")
	if ret == "" {
		ret = "SECRET_KEY"
	}

	return ret
}
